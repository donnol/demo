package main

import (
	"encoding/hex"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	expectLen := hex.EncodedLen(len(data))
	encData := hex.EncodeToString(data)
	log.Printf("%s, %d, %d\n", encData, expectLen, len(encData))

	decData, err := hex.DecodeString(encData)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v, %s\n", decData, decData)

	// 类似于 hexdump -C filename
	dumpData := hex.Dump(data)
	log.Printf("%s\n", dumpData)
}
