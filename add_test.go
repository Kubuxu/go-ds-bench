package main

import (
	"testing"

	ds "github.com/ipfs/go-datastore"
)

var mapStoreGen = func() (ds.Datastore, error) {
	return ds.NewMapDatastore(), nil
}

func BenchmarkAddTest(b *testing.B) {
	mds := ds.NewMapDatastore()
	BenchAddAt(b, mds, BenchOptions{100, 10 << 10})
}

func BenchmarkAddSeriesDefault(b *testing.B) {
	BenchAddSeriesDefault(b, mapStoreGen)
}
