package main

import "testing"

func BenchmarkAddBatchAll(b *testing.B) {
	for _, candi := range AllCandidates {
		BenchAddBatchSeriesOf(b, candi, []BenchOptions{{0, 256 << 10, 128}})
	}
}
