package main

import (
	"hash/fnv"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	hash := fnv.New128a()
	n, err := hash.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d, %d\n", n, hash.Sum(nil))
}
