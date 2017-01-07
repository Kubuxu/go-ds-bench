package main

import "testing"

func BenchmarkAddBatchSeriesDefault(b *testing.B) {
	BenchAddBatchSeriesDefault(b, CandidateMemoryMap)
}
