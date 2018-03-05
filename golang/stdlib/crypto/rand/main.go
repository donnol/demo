package main

import (
	"crypto/rand"
	"io"
	"log"
	"math/big"
)

func main() {
	for i := 0; i < 10; i++ {
		data := make([]byte, i)
		n, err := io.ReadFull(rand.Reader, data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, data)
	}

	for i := 1; i < 10; i++ {
		x := big.NewInt(int64(i))
		n, err := rand.Int(rand.Reader, x)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, x)
	}

	for i := 1; i < 10; i++ {
		p, err := rand.Prime(rand.Reader, 8)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(p)
	}

	for i := 1; i < 10; i++ {
		data := make([]byte, i)
		n, err := rand.Read(data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, data)
	}
}
