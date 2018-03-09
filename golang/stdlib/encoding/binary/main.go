package main

import (
	"encoding/binary"
	"log"
)

func main() {
	buf := []byte("Hello, I am jd")
	log.Printf("%v, %s\n", buf, buf)

	var x uint64 = 1
	r := binary.PutUvarint(buf, x)
	log.Printf("%d, %v, %s\n", r, buf, buf)

	x, r = binary.Uvarint(buf)
	log.Printf("%d, %d, %v, %s\n", x, r, buf, buf)
}
