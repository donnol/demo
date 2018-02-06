package main

import (
	"compress/flate"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	name := "about.deflate"

	writer(name)
	reader(name)
}

func reader(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fr := flate.NewReader(f)
	defer fr.Close()

	var buf = make([]byte, 32)
	n, err := fr.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, string(buf))
}

func writer(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	fw, err := flate.NewWriter(f, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer fw.Close()

	buf := []byte("hello")
	_, err = fw.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	fw.Flush() // 没有的话，读的时候报错 EOF
}
