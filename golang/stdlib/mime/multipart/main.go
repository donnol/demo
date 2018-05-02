package main

import (
	"log"
	"mime/multipart"
	"net/textproto"
)

func main() {
	fh := &multipart.FileHeader{
		Filename: "about.md",
		Header:   textproto.MIMEHeader{},
		Size:     0,
	}
	log.Println(fh)

	f, err := fh.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
}
