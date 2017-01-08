package main

import (
	"testing"

	ds "github.com/ipfs/go-datastore"
)

func BenchAddBatchAt(b *testing.B, store ds.Batching, opt BenchOptions, keys []ds.Key, bufs [][]byte) {
	//PrimeDS(b, store, opt.PrePrimeCount, opt.RecordSize)
	batch, err := store.Batch()
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		err := batch.Put(keys[i], bufs[i])
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
	err = batch.Commit()
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

		var keys []ds.Key
		var bufs [][]byte

		b.Run(newStore.Name+"/"+opt.TestDesc(), func(b *testing.B) {
			for len(keys) < b.N {
				bufs = append(bufs, RandomBuf(opt.RecordSize))
				keys = append(keys, ds.RandomKey())
			}
			BenchAddBatchAt(b, store, opt, keys, bufs)
		})
		newStore.Destroy(store)
	}
}

func BenchAddBatchSeriesDefault(b *testing.B, newStore CandidateDatastore) {
	BenchAddBatchSeriesOf(b, newStore, DefaultBenchOpts)
}
