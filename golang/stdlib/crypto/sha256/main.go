package main

import (
	"crypto/sha256"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")

	h224Data := sha256.Sum224(data)
	log.Println(h224Data)

	h256Data := sha256.Sum256(data)
	log.Println(h256Data)

	h2 := sha256.New224()
	h2.Write(data)
	h2Data := h2.Sum(nil)
	log.Println(h2Data)

	h := sha256.New()
	h.Write(data)
	hData := h.Sum(nil)
	log.Println(hData)
}
