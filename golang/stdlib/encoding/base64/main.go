package main

import (
	"encoding/base64"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	encodeData := base64.StdEncoding.EncodeToString(data)
	log.Printf("%s\n", encodeData)

	decodeData, err := base64.StdEncoding.DecodeString(encodeData)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("'%s'\n", decodeData)
}
