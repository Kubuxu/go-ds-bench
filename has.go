package main

import (
	"testing"

	ds "github.com/ipfs/go-datastore"
)

func BenchHasAt(b *testing.B, store ds.Datastore, opt BenchOptions) {
	PrimeDS(b, store, opt.PrePrimeCount, opt.RecordSize)
	buf := make([]byte, opt.RecordSize)
	keys := make([]ds.Key, b.N)
	for i := 0; i < b.N; i++ {
		buf = RandomBuf(opt.RecordSize)
		keys[i] = ds.RandomKey()

		if i%2 == 0 {
			store.Put(keys[i], buf)
		}
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := store.Has(keys[i])
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchHasSeriesOf(b *testing.B, newStore StoreGen, opts []BenchOptions) {
	for _, opt := range opts {
		store, err := newStore()
		if err != nil {
			b.Fatal(err)
		}

		b.Run(opt.TestDesc(), func(b *testing.B) {
			BenchAddAt(b, store, opt)
		})
	}
}

func BenchHasSeriesDefault(b *testing.B, newStore StoreGen) {
	BenchAddSeriesOf(b, newStore, DefaultBenchOpts)
}
