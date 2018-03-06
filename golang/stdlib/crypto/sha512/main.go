package main

import (
	"crypto/sha512"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	h512Data := sha512.Sum512(data)
	log.Println(h512Data)

	h := sha512.New()
	h.Write(data)
	hData := h.Sum(nil)
	log.Println(hData)
}
