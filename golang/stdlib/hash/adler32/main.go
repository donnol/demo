package main

import (
	"hash/adler32"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	sum := adler32.Checksum(data)
	log.Printf("%d\n", sum)

	hash := adler32.New()
	n, err := hash.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d, %d\n", n, hash.Sum32())
}
