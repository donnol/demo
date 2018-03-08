package main

import (
	"debug/gosym"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	var text uint64 = 1
	lt := gosym.NewLineTable(data, text)
	log.Printf("%+v\n", lt)

	symtab := []byte("1")
	t, err := gosym.NewTable(symtab, lt)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", t)
}
