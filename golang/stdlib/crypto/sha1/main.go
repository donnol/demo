package main

import (
	"crypto/sha1"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	hashData := sha1.Sum(data) // return [Size]byte
	log.Println(sha1.Size, hashData)

	h := sha1.New()
	h.Write(data)
	hData := h.Sum(nil) // return []byte
	log.Println(hData)
}
