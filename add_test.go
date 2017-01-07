package main

import (
	"testing"

	ds "github.com/ipfs/go-datastore"
)

func BenchmarkAddTest(b *testing.B) {
	mds := ds.NewMapDatastore()
	BenchAddAt(b, mds, BenchOptions{100, 10 << 10, 64})
}

func BenchmarkAddSeriesDefault(b *testing.B) {
	BenchAddSeriesDefault(b, CandidateMemoryMap)
}
