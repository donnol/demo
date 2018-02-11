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

type localhash struct {
	data []byte
	hash []byte
}

func (h *localhash) Write(b []byte) (int, error) {
	h.data = b
	return len(b), nil
}

func (h *localhash) Sum(b []byte) []byte {
	h.data = b
	h.hash = b
	return b
}

func (h *localhash) Reset() {
	h.data = []byte{}
	h.hash = []byte{}
}

func (h *localhash) Size() int {
	return len(h.data)
}

func (h *localhash) BlockSize() int {
	return len(h.hash)
}

func cryptoTest() {
	for i := 1; i < 10; i++ { // 不能是0, 否则会报错 panic: crypto: requested hash function #0 is unavailable
		var h = crypto.Hash(i)
		if h.Available() {
			log.Println(h.HashFunc(), h.New(), h.Size())
		}
		crypto.RegisterHash(h, func() hash.Hash {
			return &localhash{}
		})
		if h.Available() {
			log.Println(h.HashFunc(), h.New(), h.Size())
		}
	}
}
