/*
 * Copyright (C) 2017 Dgraph Labs, Inc. and Contributors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package x

import (
	"container/heap"
	"sync/atomic"

	"golang.org/x/net/trace"
)

type uint64Heap []uint64

func (u uint64Heap) Len() int               { return len(u) }
func (u uint64Heap) Less(i int, j int) bool { return u[i] < u[j] }
func (u uint64Heap) Swap(i int, j int)      { u[i], u[j] = u[j], u[i] }
func (u *uint64Heap) Push(x interface{})    { *u = append(*u, x.(uint64)) }
func (u *uint64Heap) Pop() interface{} {
	old := *u
	n := len(old)
	x := old[n-1]
	*u = old[0 : n-1]
	return x
}

// RaftValue contains the raft group and the raft proposal id.
// This is attached to the context, so the information could be passed
// down to the many posting lists, involved in mutations.
type RaftValue struct {
	Group uint32
	Index uint64
}

// mark contains raft proposal id and a done boolean. It is used to
// update the WaterMark struct about the status of a proposal.
type mark struct {
	// Either this is an (index, waiter) pair or (index, done) or (indices, done).
	index   uint64
	waiter  chan struct{}
	indices []uint64
	done    bool // Set to true if the pending mutation is done.
}

// WaterMark is used to keep track of the minimum un-finished index.  Typically, an index k becomes
// finished or "done" according to a WaterMark once Done(k) has been called
//   1. as many times as Begin(k) has, AND
//   2. a positive number of times.
//
// An index may also become "done" by calling SetDoneUntil at a time such that it is not
// inter-mingled with Begin/Done calls.
type WaterMark struct {
	Name      string
	markCh    chan mark
	doneUntil uint64
	elog      trace.EventLog
}

// Init initializes a WaterMark struct. MUST be called before using it.
func (w *WaterMark) Init() {
	w.markCh = make(chan mark, 10000)
	w.elog = trace.NewEventLog("Watermark", w.Name)
	go w.process()
}

func (w *WaterMark) Begin(index uint64) {
	w.markCh <- mark{index: index, done: false}
}
func (w *WaterMark) BeginMany(indices []uint64) {
	w.markCh <- mark{index: 0, indices: indices, done: false}
}

func (w *WaterMark) Done(index uint64) {
	w.markCh <- mark{index: index, done: true}
}
func (w *WaterMark) DoneMany(indices []uint64) {
	w.markCh <- mark{index: 0, indices: indices, done: true}
}

// DoneUntil returns the maximum index until which all tasks are done.
func (w *WaterMark) DoneUntil() uint64 {
	return atomic.LoadUint64(&w.doneUntil)
}

func (w *WaterMark) SetDoneUntil(val uint64) {
	atomic.StoreUint64(&w.doneUntil, val)
}

func (w *WaterMark) WaitForMark(index uint64) {
	if w.DoneUntil() >= index {
		return
	}
	waitCh := make(chan struct{})
	w.markCh <- mark{index: index, waiter: waitCh}
	<-waitCh
	return
}

// process is used to process the Mark channel. This is not thread-safe,
// so only run one goroutine for process. One is sufficient, because
// all goroutine ops use purely memory and cpu.
func (w *WaterMark) process() {
	var indices, waiterIndices uint64Heap
	// pending maps raft proposal index to the number of pending mutations for this proposal.
	pending := make(map[uint64]int)
	waiters := make(map[uint64][]chan struct{})

	heap.Init(&indices)
	heap.Init(&waiterIndices)
	var loop uint64

	processOne := func(index uint64, done bool) {
		// If not already done, then set. Otherwise, don't undo a done entry.
		prev, present := pending[index]
		if !present {
			heap.Push(&indices, index)
		}

		delta := 1
		if done {
			delta = -1
		}
		pending[index] = prev + delta

		loop++
		if len(indices) > 0 && loop%10000 == 0 {
			min := indices[0]
			w.elog.Printf("WaterMark %s: Done entry %4d. Size: %4d Watermark: %-4d Looking for: %-4d. Value: %d\n",
				w.Name, index, len(indices), w.DoneUntil(), min, pending[min])
		}

		// Update mark by going through all indices in order; and checking if they have
		// been done. Stop at the first index, which isn't done.
		doneUntil := w.DoneUntil()
		AssertTrue(doneUntil < index)

		until := doneUntil
		loops := 0

		for len(indices) > 0 {
			min := indices[0]
			if done := pending[min]; done != 0 {
				break // len(indices) will be > 0.
			}
			heap.Pop(&indices)
			delete(pending, min)
			until = min
			loops++
		}
		if until != doneUntil {
			for len(waiterIndices) > 0 {
				min := waiterIndices[0]
				if min > until {
					break
				}
				// Partly duplicated with what's above.
				heap.Pop(&waiterIndices)
				toNotify := waiters[min]
				for _, ch := range toNotify {
					close(ch)
				}
				delete(waiters, min)
			}

			AssertTrue(atomic.CompareAndSwapUint64(&w.doneUntil, doneUntil, until))
			w.elog.Printf("%s: Done until %d. Loops: %d\n", w.Name, until, loops)
		}
	}

	for mark := range w.markCh {
		if mark.waiter != nil {
			doneUntil := atomic.LoadUint64(&w.doneUntil)
			if doneUntil >= mark.index {
				close(mark.waiter)
			} else {
				ws, ok := waiters[mark.index]
				if !ok {
					heap.Push(&waiterIndices, mark.index)
					waiters[mark.index] = []chan struct{}{mark.waiter}
				} else {
					waiters[mark.index] = append(ws, mark.waiter)
				}
			}
		} else {
			if mark.index > 0 {
				processOne(mark.index, mark.done)
			}
			for _, index := range mark.indices {
				processOne(index, mark.done)
			}
		}
	}
}
