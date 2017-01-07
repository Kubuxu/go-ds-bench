package main

import (
	"fmt"
	"math/rand"
)

var (
	preRandom []byte
	pos       int
)

func init() {
	preRandom = make([]byte, 8<<20)
	_, err := rand.Read(preRandom)
	if err != nil {
		panic(err)
	}

	pos = 0
}

func RandomBuf(req int) []byte {
	if req > len(preRandom) {
		fmt.Printf("req %d, len %d\n", req, len(preRandom))
		panic("aka rand: requested len too long")
	}

	if req+pos >= len(preRandom) {
		pos = 0
	}

	defer func() {
		pos++
	}()

	return preRandom[pos : pos+req]
}
