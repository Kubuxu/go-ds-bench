package main

import (
	ds "github.com/ipfs/go-datastore"
)

var CandidateMemoryMap = CandidateDatastore{
	Name: "memory-map",
	Create: func() (ds.Datastore, error) {
		return ds.NewMapDatastore(), nil
	},
	Destroy: func(ds.Datastore) {},
}

type CandidateDatastore struct {
	Name    string
	Create  func() (ds.Datastore, error)
	Destroy func(ds.Datastore)
}
