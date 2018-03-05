package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"math/big"
)

func main() {
	c := elliptic.P256()

	priv, x, y, err := elliptic.GenerateKey(c, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s, %d, %d", priv, x, y)

	data := elliptic.Marshal(c, x, y)
	log.Printf("%s", data)

	x, y = elliptic.Unmarshal(c, data)
	log.Printf("%d, %d", x, y)

	ec := &elliptic.CurveParams{
		P:    big.NewInt(1),
		N:    big.NewInt(-1),
		B:    big.NewInt(0),
		Gx:   big.NewInt(1),
		Gy:   big.NewInt(1),
		Name: "hello",
	}
	log.Printf("%+v", ec)
	x, y = ec.Add(big.NewInt(1), big.NewInt(1), big.NewInt(2), big.NewInt(2))
	log.Printf("%+v, %d, %d", ec, x, y)
}
