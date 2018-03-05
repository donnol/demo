package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	priv, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(priv)

	// 公钥加密
	data := []byte("Hello, I am jd")
	encryptData, err := rsa.EncryptPKCS1v15(rand.Reader, &priv.PublicKey, data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(encryptData, string(encryptData))

	// 私钥解密
	decryptData, err := rsa.DecryptPKCS1v15(rand.Reader, priv, encryptData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(decryptData, string(decryptData))

	// 私钥签名
	sha := sha256.New()
	sha.Write(data)
	hashData := sha.Sum(nil)
	signData, err := rsa.SignPKCS1v15(rand.Reader, priv, crypto.SHA256, hashData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(signData)

	// 公钥验证
	err = rsa.VerifyPKCS1v15(&priv.PublicKey, crypto.SHA256, hashData, signData)
	if err != nil {
		log.Fatal(err)
	}
}
