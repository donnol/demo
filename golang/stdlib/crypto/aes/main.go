package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// 高级加密标准
	// 对称密钥加密
	// 代换-置换网络
	// 快速地加解密
	// 只需要很少的内存
	// 针对AES唯一的成功攻击是旁道攻击
	aesTest()

	key := []byte("LKHlhb899Y09olUi")
	for _, tc := range []string{
		"Hello World",
		"I am Mr.Lau",
		"This is a very long statement",
		"This is a very very very very very very very very very long statement",
	} {
		encryptMsg, _ := encrypt(key, tc)
		fmt.Println(encryptMsg) // skkzXCbis0E3V5tEEzCCHJY0MP1wUk-yr_PHm1COns0
		msg, _ := decrypt(key, encryptMsg)
		fmt.Println(msg) // Hello World
	}
}

func aesTest() {
	log.Println(aes.BlockSize)

	key := []byte("key1")
	key = bytes.Repeat(key, 8)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(block.BlockSize())

	src := []byte("hello")
	src = bytes.Repeat(src, 6)
	dst := make([]byte, len(src))
	block.Encrypt(dst, src)
	// panic: crypto/aes: input not full block -- 长度不能小于 aes.BlockSize
	// panic: crypto/aes: output not full block -- 长度不能小于 aes.BlockSize
	log.Println(src, "after aes encrypt", dst)

	src2 := make([]byte, len(dst))
	block.Decrypt(src2, dst)
	log.Println(dst, "after aes decrypt", src2)
}

// from https://gist.github.com/stupidbodo/601b68bfef3449d1b8d9
func addBase64Padding(value string) string {
	m := len(value) % 4
	if m != 0 {
		value += strings.Repeat("=", 4-m)
	}

	return value
}

func removeBase64Padding(value string) string {
	return strings.Replace(value, "=", "", -1)
}

// Pad 垫
func Pad(src []byte) []byte {
	padding := aes.BlockSize - len(src)%aes.BlockSize // 缺 padding 字节，不够一个 block
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...) // 将 padding 长度的 padding 追加到源字节切片后
}

// Unpad 去垫
func Unpad(src []byte) ([]byte, error) {
	length := len(src)
	if length == 0 {
		return nil, errors.New("src is empty")
	}
	unpadding := int(src[length-1]) // 最后一个字节是之前垫的字节，也是垫的字节长度

	if unpadding > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - unpadding)], nil // 源串
}

func encrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	msg := Pad([]byte(text)) // msg 的长度必须是 ase.BlockSize 倍
	ciphertext := make([]byte, aes.BlockSize+len(msg))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil { // 随机填充字节
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
	finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
	return finalMsg, nil
}

func decrypt(key []byte, text string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
	if err != nil {
		return "", err
	}

	if (len(decodedMsg) % aes.BlockSize) != 0 {
		return "", errors.New("blocksize must be multipe of decoded message length")
	}

	iv := decodedMsg[:aes.BlockSize]
	msg := decodedMsg[aes.BlockSize:]

	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(msg, msg)

	unpadMsg, err := Unpad(msg)
	if err != nil {
		return "", err
	}

	return string(unpadMsg), nil
}
