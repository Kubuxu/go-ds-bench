package main

import (
	ds "github.com/ipfs/go-datastore"
)

var CandidateMemoryMap = CandidateDatastore{
	Name: "memory-map",
	Create: func() (ds.Batching, error) {
		return ds.NewMapDatastore(), nil
	},
	Destroy: func(ds.Batching) {},
}

type CandidateDatastore struct {
	Name    string
	Create  func() (ds.Batching, error)
	Destroy func(ds.Batching)
}
