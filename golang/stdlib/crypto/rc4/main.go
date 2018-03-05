package main

import (
	"crypto/rc4"
	"log"
)

func main() {
	key := []byte("Hello, I am jd")

	// 加密
	cipher, err := rc4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	src := []byte("origin data")
	// dst := make([]byte, len(src))
	cipher.XORKeyStream(src, src)
	log.Println(src, string(src))

	// cipher.Reset()

	// 解密--使用相同 key 重新 new 一个 cipher
	cipher, err = rc4.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	cipher.XORKeyStream(src, src)
	log.Println(src, string(src))
}
