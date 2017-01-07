package main

import "testing"

func BenchmarkGetSeriesDefault(b *testing.B) {
	BenchGetSeriesDefault(b, CandidateMemoryMap)
}
