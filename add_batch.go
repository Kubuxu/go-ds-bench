package main

import (
	"testing"

	ds "github.com/ipfs/go-datastore"
)

func BenchAddBatchAt(b *testing.B, store ds.Batching, opt BenchOptions) {
	PrimeDS(b, store, opt.PrePrimeCount, opt.RecordSize)
	bufs := make([][]byte, b.N)
	keys := make([]ds.Key, b.N)
	for i := 0; i < b.N; i++ {
		bufs[i] = RandomBuf(opt.RecordSize)
		keys[i] = ds.RandomKey()
	}

	b.ResetTimer()
	batch, err := store.Batch()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		batch.Put(keys[i], bufs[i])
		if err != nil {
			b.Fatal(err)
		}

		if i%opt.BatchSize == opt.BatchSize-1 {
			err = batch.Commit()
			if err != nil {
				b.Fatal(err)
			}
			batch, err = store.Batch()
			if err != nil {
				b.Fatal(err)
			}
		}
	}
	batch, err = store.Batch()
	if err != nil {
		b.Fatal(err)
	}

}

func BenchAddBatchSeriesOf(b *testing.B, newStore CandidateDatastore, opts []BenchOptions) {
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

func BenchAddBatchSeriesDefault(b *testing.B, newStore CandidateDatastore) {
	BenchAddBatchSeriesOf(b, newStore, DefaultBenchOpts)
}
