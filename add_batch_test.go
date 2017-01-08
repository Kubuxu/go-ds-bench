package main

import "testing"

func BenchmarkAddBatchSeriesDefault(b *testing.B) {
	for _, c := range AllCandidates {
		BenchAddBatchSeriesDefault(b, c)
	}
}
