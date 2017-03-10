package main

import (
	"io/ioutil"
	"os"

	fsbs "github.com/ipfs/fsbs"
	ds "github.com/ipfs/go-datastore"
	flatfs "github.com/ipfs/go-ds-flatfs"
	bolt "github.com/whyrusleeping/bolt-datastore"
)

func emptyDtor(ds.Batching) {
}

var AllCandidates = []CandidateDatastore{
	CandidateMemoryMap,
	CandidateFsbs,
	CandidateBolt,
	CandidateFlatfs,
	CandidateFlatfsNoSync,
}

var CandidateMemoryMap = CandidateDatastore{
	Name: "memory-map",
	Create: func() (ds.Batching, error) {
		return ds.NewMapDatastore(), nil
	},
	Destroy: emptyDtor,
}

var CandidateBolt = CandidateDatastore{
	Name: "bolt",
	Create: func() (ds.Batching, error) {
		os.Mkdir("bolt", 0775)

		dir, err := ioutil.TempDir("bolt", "bench")
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(dir, 0775)
		if err != nil {
			return nil, err
		}

		return bolt.NewBoltDatastore(dir, "foo")
	},
	Destroy: func(b ds.Batching) {
		dbpath := b.(*bolt.BoltDatastore).Path
		os.Remove(dbpath)
	},
}

var CandidateFsbs = CandidateDatastore{
	Name: "fsbs",
	Create: func() (ds.Batching, error) {

		os.Mkdir("fsbs", 0775)
		dir, err := ioutil.TempDir("fsbs", "bench")
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(dir, 0775)
		if err != nil {
			return nil, err
		}

		return fsbs.NewFsbsDS(dir)

	},

	Destroy: func(b ds.Batching) {
		os.Remove(b.(*fsbs.Fsbsds).Path)
	},
}

func flatfsCtor(sync bool) func() (ds.Batching, error) {
	return func() (ds.Batching, error) {
		os.Mkdir("flatfs", 0775)

		dir, err := ioutil.TempDir("flatfs", "bench")
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(dir, 0775)
		if err != nil {
			return nil, err
		}

		return flatfs.CreateOrOpen(dir, flatfs.NextToLast(2), sync)
	}
}

var CandidateFlatfs = CandidateDatastore{
	Name:    "flatfs",
	Create:  flatfsCtor(true),
	Destroy: emptyDtor,
}

var CandidateFlatfsNoSync = CandidateDatastore{
	Name:    "flatfs-nosync",
	Create:  flatfsCtor(false),
	Destroy: emptyDtor,
}

type CandidateDatastore struct {
	Name    string
	Create  func() (ds.Batching, error)
	Destroy func(ds.Batching)
}
