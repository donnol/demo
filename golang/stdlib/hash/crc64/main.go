package main

import (
	"hash/crc64"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	tab := crc64.MakeTable(0x123456)
	sum := crc64.Checksum(data, tab)
	log.Printf("%d\n", sum)

	p := []byte("suffix")
	sum = crc64.Update(sum, tab, p)
	log.Printf("%d\n", sum)
}
