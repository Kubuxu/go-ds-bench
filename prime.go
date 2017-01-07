package main

import (
	"math/rand"
	"testing"

	ds "github.com/ipfs/go-datastore"
)

func PrimeDS(tb testing.TB, store ds.Datastore, count, blockSize int) {

	buf := make([]byte, blockSize)
	for i := 0; i < count; i++ {
		_, err := rand.Read(buf)
		if err != nil {
			tb.Fatal(err)
		}
		store.Put(ds.RandomKey(), buf)
	}

}
