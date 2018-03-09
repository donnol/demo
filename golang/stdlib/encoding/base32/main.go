package main

import (
	"bytes"
	"encoding/base32"
	"log"
)

func main() {
	var w = new(bytes.Buffer)
	wc := base32.NewEncoder(base32.HexEncoding, w)

	data := []byte("Hello, I am jd")
	n, err := wc.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	wc.Close()

	log.Printf("%d, %d, '%s'\n", n, w.Len(), data)

	dst := make([]byte, w.Len())
	n, err = w.Read(dst)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%d, '%s'\n", n, dst)

	var w2 = bytes.NewBuffer(dst)
	r := base32.NewDecoder(base32.HexEncoding, w2)
	originData := make([]byte, len(data)+1)
	n, err = r.Read(originData)
	if err != nil {
		log.Fatal(err)
	}
	// 当 originData 的长度设置为 len(data) 时，输出 10, 14, 'Hello, I a    '
	// 当 originData 的长度设置为 len(data)+1 时，输出 10, 15, 'Hello, I am jd '，后面会多个空格
	log.Printf("%d, %d, '%s'\n", n, len(originData), originData)
}
