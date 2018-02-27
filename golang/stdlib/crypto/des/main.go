package main

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"io"
	"log"
)

func main() {
	log.Println(des.BlockSize)

	// 生成 key
	key := make([]byte, des.BlockSize, des.BlockSize)
	_, err := io.ReadFull(rand.Reader, key)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(key)

	// 数据
	text := []byte("Hello, I am Mr.right")
	src := make([]byte, des.BlockSize+len(text)) // 前 des.BlockSize 个字节填充随机字符
	iv := src[:des.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		log.Fatal(err)
	}

	// 加密
	block, err := des.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(src[block.BlockSize():], text)
	log.Println(src)

	// 解密
	block, err = des.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	iv = src[:block.BlockSize()]
	stream = cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(src[block.BlockSize():], src[block.BlockSize():])
	log.Println(string(src[block.BlockSize():]))
}
