package main

import "testing"

func BenchmarkHasSeriesDefault(b *testing.B) {
	BenchHasSeriesDefault(b, mapStoreGen)
}
