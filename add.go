package main

import (
	"testing"

	ds "github.com/ipfs/go-datastore"
)

type StoreGen func() (ds.Batching, error)

func BenchAddAt(b *testing.B, store ds.Batching, opt BenchOptions) {
	PrimeDS(b, store, opt.PrePrimeCount, opt.RecordSize)
	bufs := make([][]byte, b.N)
	keys := make([]ds.Key, b.N)
	for i := 0; i < b.N; i++ {
		bufs[i] = RandomBuf(opt.RecordSize)
		keys[i] = ds.RandomKey()
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := store.Put(keys[i], bufs[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchAddSeriesOf(b *testing.B, newStore CandidateDatastore, opts []BenchOptions) {
	for _, opt := range opts {
		store, err := newStore.Create()
		if err != nil {
			b.Fatal(err)
		}

		b.Run(opt.TestDesc(), func(b *testing.B) {
			BenchAddAt(b, store, opt)
		})
		newStore.Destroy(store)
	}
}

func BenchAddSeriesDefault(b *testing.B, newStore CandidateDatastore) {
	BenchAddSeriesOf(b, newStore, DefaultBenchOpts)
}
