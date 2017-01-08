package main

import "testing"

func BenchmarkAddBatchSeriesDefault(b *testing.B) {
	BenchAddBatchSeriesDefault(b, CandidateBolt)
	BenchAddBatchSeriesDefault(b, CandidateFlatfs)
	BenchAddBatchSeriesDefault(b, CandidateFlatfsNoSync)
	BenchAddBatchSeriesDefault(b, CandidateMemoryMap)
}
