package main

import (
	"bytes"
	"index/suffixarray"
	"log"
	"os"
	"regexp"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	data := []byte("Hello, I am jd.")
	index := suffixarray.New(data)
	log.Printf("%s\n", index.Bytes())
	reg, err := regexp.Compile("^a$")
	if err != nil {
		log.Fatal(err)
	}
	res := index.FindAllIndex(reg, -1)
	log.Printf("%v\n", res)

	sub := []byte("am")
	res2 := index.Lookup(sub, -1)
	log.Printf("%v\n", res2)

	buf := bytes.NewBuffer(data)
	err = index.Read(buf) // 传入文件句柄时 panic: runtime error: makeslice: len out of range
	if err != nil {
		log.Fatal(err) // 传入bytes.Buffer时 unexpected EOF
	}

	f2, err := os.OpenFile("copy.md", os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModePerm)
	err = index.Write(f2)
	if err != nil {
		log.Fatal(err)
	}
}
