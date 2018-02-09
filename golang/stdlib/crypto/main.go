package main

import (
	"crypto"
	"hash"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cryptoTest()
}

type localhash struct{}

func (h localhash) Write(b []byte) (int, error) {
	return len(b), nil
}
func (h localhash) Sum(b []byte) []byte {
	return b
}
func (h localhash) Reset() {
}
func (h localhash) Size() int {
	return 0
}
func (h localhash) BlockSize() int {
	return 0
}

func cryptoTest() {
	for i := 0; i < 10; i++ {
		var h = crypto.Hash(i)
		// log.Println(h.Available(), h.HashFunc(), h.New(), h.Size())
		crypto.RegisterHash(h, func() hash.Hash {
			return localhash{}
		})
		// log.Println(h.Available(), h.HashFunc(), h.New(), h.Size())
	}
}
