package main

import (
	"compress/gzip"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	name := "about.gzip"
	writer(name)
	reader(name)
}

func reader(name string) {
	fi, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gr, err := gzip.NewReader(f)
	if err != nil {
		log.Fatal(err)
	}
	defer gr.Close()

	var buf = make([]byte, fi.Size())
	n, err := gr.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, fi.Size(), "'"+string(buf)+"'", gr.Header)
}

func writer(name string) {
	f, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	gw := gzip.NewWriter(f)
	defer gw.Close()

	buf := []byte("hello")
	_, err = gw.Write(buf)
	if err != nil {
		log.Fatal(err)
	}
	gw.Flush()
}
