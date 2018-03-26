package main

import (
	"hash"
)

type myHash struct{}

func (m myHash) Write(p []byte) (n int, err error) {
	return
}

func (m myHash) Sum(b []byte) (p []byte) {
	return
}

func (m myHash) Reset() {

}

func (m myHash) Size() (n int) {
	return
}

func (m myHash) BlockSize() (n int) {
	return
}

type myHash32 struct {
	myHash
}

func (m myHash32) Sum32() (sum uint32) {
	return
}

type myHash64 struct {
	myHash
}

func (m myHash64) Sum64() (sum uint64) {
	return
}

func main() {
	var _ hash.Hash = myHash{}
	var _ hash.Hash32 = myHash32{}
	var _ hash.Hash64 = myHash64{}
}
