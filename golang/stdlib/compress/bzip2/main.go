package main

import (
	"compress/bzip2"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	filename := "about.md"
	reader(filename)
}

func reader(filename string) {
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	br := bzip2.NewReader(f)
	var buf []byte
	// 当打开的文件不是 bzip2 格式时，报错
	// 	bzip2 data invalid: bad magic value
	_, err = br.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
}
