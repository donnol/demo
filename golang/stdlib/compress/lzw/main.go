package main

import (
	"compress/lzw"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	name := "about.lzw"
	writer(name)
	reader(name)
}

func reader(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	lr := lzw.NewReader(f, lzw.LSB, 8)
	defer lr.Close()

	var buf = make([]byte, 1<<5)
	n, err := lr.Read(buf)
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

	lw := lzw.NewWriter(f, lzw.LSB, 8)
	defer lw.Close()

	buf := []byte("hello")
	_, err = lw.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
}
