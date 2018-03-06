package main

import (
	"crypto/rand"
	"crypto/x509"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile)

	data := []byte("Hello, I am jd")
	parent, err := x509.ParseCertificate(data)
	if err != nil {
		log.Fatal(err)
	}
	pub := 1
	priv := 2
	cert, err := x509.CreateCertificate(rand.Reader, parent, parent, pub, priv)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(cert)

	// https://golang.org/src/crypto/x509/example_test.go
}
