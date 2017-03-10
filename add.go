package main

import (
	"syscall"
	"testing"

	ds "github.com/ipfs/go-datastore"
)

type StoreGen func() (ds.Batching, error)

func BenchAddAt(b *testing.B, store ds.Batching, opt BenchOptions, keys []ds.Key, bufs [][]byte) {
	//PrimeDS(b, store, opt.PrePrimeCount, opt.RecordSize)
	b.SetBytes(int64(opt.RecordSize))
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

		var keys []ds.Key
		var bufs [][]byte

		b.Run(newStore.Name+"/"+opt.TestDesc(), func(b *testing.B) {
			for len(keys) < b.N {
				bufs = append(bufs, RandomBuf(opt.RecordSize))
				keys = append(keys, ds.RandomKey())
			}
			syscall.Sync()
			BenchAddAt(b, store, opt, keys, bufs)
		})
		newStore.Destroy(store)
	}
}

func BenchAddSeriesDefault(b *testing.B, newStore CandidateDatastore) {
	BenchAddSeriesOf(b, newStore, DefaultBenchOpts)
}
