package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
)

func main() {
	c := elliptic.P256()
	priv, err := ecdsa.GenerateKey(c, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// 签名
	hash := []byte("Hello, please use ecdsa sign me")
	r, s, err := ecdsa.Sign(rand.Reader, priv, hash)
	if err != nil {
		log.Fatal(err)
	}

	// 验证
	if !ecdsa.Verify(&priv.PublicKey, hash, r, s) {
		log.Fatal("verify failed")
	}
}
