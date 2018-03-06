package main

import (
	"crypto/rand"
	"crypto/tls"
	"log"
)

func main() {
	config := &tls.Config{
		Rand:               rand.Reader,
		InsecureSkipVerify: true, // 不检查服务器证书
	}
	conn, err := tls.Dial("tcp", ":8825", config)
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 20)
	_, err = conn.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(data, "'", string(data), "'")
}
