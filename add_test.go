package main

import "testing"

func BenchmarkAddTest(b *testing.B) {
	//mds := ds.NewMapDatastore()
	//	BenchAddAt(b, mds, BenchOptions{100, 10 << 10, 64})
}

func BenchmarkAddSeriesDefault(b *testing.B) {
	for _, c := range AllCandidates {
		BenchAddSeriesDefault(b, c)
	}
}
