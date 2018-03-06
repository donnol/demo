package main

import (
	"crypto/x509/pkix"
	"encoding/asn1"
	"log"
)

func main() {
	var alg pkix.AlgorithmIdentifier
	alg.Algorithm = asn1.ObjectIdentifier{1}
	alg.Parameters = asn1.RawValue{
		Class:      1,
		Tag:        2,
		IsCompound: true,
	}
	log.Println(alg)
}
