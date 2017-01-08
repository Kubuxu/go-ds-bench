package main

import (
	"math/rand"
	"testing"

	ds "github.com/ipfs/go-datastore"
)

const primeMaxBatchSize = 1 << 30 // 1 GiB

func PrimeDS(tb testing.TB, store ds.Batching, count, blockSize int) {
	buf := make([]byte, blockSize)
	maxBatchCount := primeMaxBatchSize / blockSize

	b, err := store.Batch()
	if err != nil {
		tb.Fatal(err)
	}

	for i := 0; i < count; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			tb.Fatal(err)
		}
		err = b.Put(ds.RandomKey(), buf)
		if err != nil {
			tb.Fatal(err)
		}

		if i%maxBatchCount == maxBatchCount-1 {
			err = b.Commit()
			if err != nil {
				tb.Fatal(err)
			}
			b, err = store.Batch()
			if err != nil {
				tb.Fatal(err)
			}
		}
	}

	err = b.Commit()
	if err != nil {
		tb.Fatal(err)
	}
}
