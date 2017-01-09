package main

import "testing"

func BenchmarkAddBatchAll(b *testing.B) {
	for _, candi := range AllCandidates {
		BenchAddBatchSeriesOf(b, candi, []BenchOptions{{0, 10 << 10, 128}})
	}
}
