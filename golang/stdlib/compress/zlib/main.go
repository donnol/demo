package main

import (
	"compress/zlib"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	name := "about.zlib"
	writer(name)
	reader(name)
}

func reader(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	zr, err := zlib.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	var buf = make([]byte, 1<<5)
	n, err := zr.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, "'"+string(buf)+"'")
}

func writer(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	zw := zlib.NewWriter(f)
	defer zw.Close()

	buf := []byte("hello")
	_, err = zw.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	zw.Flush()
}
