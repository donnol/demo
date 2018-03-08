package main

import (
	"debug/plan9obj"
	"log"
)

func main() {
	// file, err := plan9obj.Open("main.go") // bad magic number '[112 97 99 107]' in record at byte 0x0
	file, err := plan9obj.Open("hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", file)

	// {
	// 	Magic:35479
	// 	Bss:126272
	// 	Entry:2405136
	// 	PtrSize:8
	// 	LoadAddress:2097152
	// 	HdrSize:40
	// }
	log.Printf("%+v\n", file.FileHeader)
}
