package main

import (
	"bytes"
	"encoding/pem"
	"log"
)

func main() {
	block := &pem.Block{
		Type: "RSA PRIVATE KEY",
		Headers: map[string]string{
			"Animal": "Gopher",
		},
		Bytes: []byte("Hello, I am jd"),
	}
	var w = new(bytes.Buffer)
	err := pem.Encode(w, block)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", w.String())

	p, rest := pem.Decode(w.Bytes())
	log.Printf("%+v, %v\n", p, rest)
}
