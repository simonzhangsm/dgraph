package main

import (
	"encoding/binary"
	"hash/crc64"
	"io/ioutil"
	"os"
	"sort"
	"sync/atomic"

	"github.com/dgraph-io/badger"
	"github.com/dgraph-io/dgraph/bp128"
	"github.com/dgraph-io/dgraph/protos"
	"github.com/dgraph-io/dgraph/x"
)

func newPlBuilder(tmpDir string, prog *progress) *plBuilder {
	badgerDir, err := ioutil.TempDir(tmpDir, "dgraph_bulkloader")
	x.Check(err)
	kv, err := defaultBadger(badgerDir)
	x.Check(err)
	return &plBuilder{NewKVWriter(kv), badgerDir, prog}
}

type plBuilder struct {
	kvw       *KVWriter
	badgerDir string
	prog      *progress
}

func (b *plBuilder) cleanUp() {
	// Don't need any persistence, but still Close() anyway to close all FDs
	// before nuking the data directory.
	x.Check(b.kvw.kv.Close())
	x.Check(os.RemoveAll(b.badgerDir))
}

func (b *plBuilder) addPosting(postingListKey []byte, posting *protos.Posting, countGroupHash uint64) {

	atomic.AddInt64(&b.prog.tmpKeyTotal, 1)

	plKeyHash := crc64.Checksum(postingListKey, crc64.MakeTable(crc64.ISO))
	var kBuf [24]byte
	binary.BigEndian.PutUint64(kBuf[:], countGroupHash) // Byte order doesn't matter (it's for grouping).
	binary.BigEndian.PutUint64(kBuf[8:], plKeyHash)     // Byte order doesn't matter (it's for grouping).
	binary.BigEndian.PutUint64(kBuf[16:], posting.Uid)  // Byte order is important for iteration order.

	// TODO: We could save 4 bytes for UID postings by omitting the posting
	// list length.

	// TODO: Another space saving could be to store PLKs separately, i.e. in
	// another set of keys #(PLK) => PLK. Could use the meta byte to
	// differentiate.

	if posting.GetPostingType() == protos.Posting_REF {
		vBuf := make([]byte, 4+len(postingListKey))
		binary.BigEndian.PutUint32(vBuf, uint32(len(postingListKey)))
		copy(vBuf[4:], postingListKey)
		b.kvw.Set(kBuf[:], vBuf, 0x01)
	} else {
		postingBuf, err := posting.Marshal()
		x.Check(err)
		vBuf := make([]byte, 4+len(postingListKey)+len(postingBuf))
		binary.BigEndian.PutUint32(vBuf, uint32(len(postingListKey)))
		copy(vBuf[4:], postingListKey)
		copy(vBuf[4+len(postingListKey):], postingBuf)
		b.kvw.Set(kBuf[:], vBuf, 0x00)
	}
}

func (b *plBuilder) buildPostingLists(target *badger.KV, ss schemaStore) {

	b.kvw.Wait()

	batchTarget := NewKVWriter(target)

	counts := map[int][]uint64{}

	pl := &protos.PostingList{}
	uids := []uint64{}
	iter := b.kvw.kv.NewIterator(badger.DefaultIteratorOptions)
	iter.Rewind()
	if !iter.Valid() {
		// There were no posting lists to build.
		return
	}
	k := unpackPostingListKey(iter.Item())
	kHash := unpackPostingListKeyHash(iter.Item())
	for iter.Valid() {

		// Add to PL
		uids = append(uids, unpackUidPosting(iter.Item()))
		if iter.Item().UserMeta() == 0x00 {
			p := unpackFullPosting(iter.Item())
			pl.Postings = append(pl.Postings, p)
		}

		countGroupHash := unpackCountGroupHash(iter.Item())

		// Determine if we're at the end of a single posting list.
		finalise := false
		iter.Next()
		atomic.AddInt64(&b.prog.tmpKeyProg, 1)
		var newK []byte
		var newKHash uint64
		if iter.Valid() {
			newK = unpackPostingListKey(iter.Item())
			newKHash = unpackPostingListKeyHash(iter.Item())
			if kHash != newKHash {
				finalise = true
			}
		} else {
			finalise = true
		}

		// Write posting list out to target.
		if finalise {

			// If we saw any full postings, then use a proto.PostingList as the
			// value. But include the UID-only postings in the posting list
			// (not just the proto.Posting values).

			useFullPostings := len(pl.Postings) > 0

			if useFullPostings {
				pl.Uids = bp128.DeltaPack(uids)
				plBuf, err := pl.Marshal()
				x.Check(err)
				batchTarget.Set(k, plBuf, 0x00)
			} else {
				batchTarget.Set(k, bp128.DeltaPack(uids), 0x01)
			}

			parsedK := x.Parse(k)
			if (parsedK.IsData() || parsedK.IsReverse()) && ss.m[parsedK.Attr].GetCount() {
				cnt := len(uids)
				counts[cnt] = append(counts[cnt], parsedK.Uid)
			}

			// Reset for next posting list.
			pl.Postings = nil
			pl.Uids = nil
			uids = nil

			// Create PLs for count index.
			if !iter.Valid() || countGroupHash != unpackCountGroupHash(iter.Item()) {
				for cnt, uids := range counts {
					// TODO: Is sort.Slice slow due to reflection? If so, use a faster sort method.
					sort.Slice(uids, func(i, j int) bool { return uids[i] < uids[j] })
					key := x.CountKey(parsedK.Attr, uint32(cnt), parsedK.IsReverse())
					val := bp128.DeltaPack(uids)
					batchTarget.Set(key, val, 0x01)
				}
				counts = map[int][]uint64{} // TODO: Possibly faster to clear map while iterating. Profile to work out.
			}
		}

		k = newK
		kHash = newKHash
	}

	batchTarget.Wait()
}

func unpackPostingListKey(item *badger.KVItem) []byte {
	v := item.Value()
	plKeyLen := binary.BigEndian.Uint32(v)
	plKey := make([]byte, plKeyLen)
	copy(plKey, v[4:4+plKeyLen])
	return plKey
}

func unpackPostingListKeyHash(item *badger.KVItem) uint64 {
	return binary.BigEndian.Uint64(item.Key()[8:])
}

// unpacks a full posting into the posting list key and posting.
func unpackFullPosting(item *badger.KVItem) *protos.Posting {
	v := item.Value()
	plKeyLen := binary.BigEndian.Uint32(v)
	posting := new(protos.Posting)
	x.Check(posting.Unmarshal(v[4+plKeyLen:]))
	return posting
}

// unpacks a UID posting into the posting list key and uid.
func unpackUidPosting(item *badger.KVItem) uint64 {
	return binary.BigEndian.Uint64(item.Key()[16:])
}

func unpackCountGroupHash(item *badger.KVItem) uint64 {
	return binary.BigEndian.Uint64(item.Key())
}