package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
	"github.com/gogo/protobuf/proto"
)

func writeDenormalisedPostings(dir string, postingsIn <-chan *protos.DenormalisedPosting, prog *progress) {

	var fileNum int
	var postings []*protos.DenormalisedPosting
	var wg sync.WaitGroup
	var sz int

	processBatch := func() {
		wg.Add(1)
		filename := filepath.Join(dir, fmt.Sprintf("%06d.bin", fileNum))
		fileNum++
		ps := postings
		go func() {
			sortAndDump(filename, ps, prog)
			wg.Done()
		}()
		postings = nil
		sz = 0
	}

	for posting := range postingsIn {
		postings = append(postings, posting)
		sz += posting.Size()
		if sz > 32<<20 {
			processBatch()
		}
	}
	if len(postings) > 0 {
		processBatch()
	}

	wg.Wait()
}

type dpl []*protos.DenormalisedPosting

func (d dpl) Less(i, j int) bool { return bytes.Compare(d[i].PostingListKey, d[j].PostingListKey) < 0 }
func (d dpl) Len() int           { return len(d) }
func (d dpl) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

func sortAndDump(filename string, postings []*protos.DenormalisedPosting, prog *progress) {

	atomic.AddInt64(&prog.sorting, 1)
	//sort.Slice(postings, func(i, j int) bool {
	//return bytes.Compare(postings[i].PostingListKey, postings[j].PostingListKey) < 0
	//})
	sort.Sort(dpl(postings))
	atomic.AddInt64(&prog.sorting, -1)

	atomic.AddInt64(&prog.writing, 1)
	var buf proto.Buffer
	for _, posting := range postings {
		x.Check(buf.EncodeMessage(posting))
	}

	fd, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	x.Checkf(err, "Could not open tmp file.")
	x.Check2(fd.Write(buf.Bytes()))
	x.Check(fd.Close())
	atomic.AddInt64(&prog.writing, -1)
}
