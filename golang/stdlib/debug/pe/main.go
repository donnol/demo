package main

import (
	"debug/pe"
	"log"
)

func main() {
	// file, err := pe.Open("main.go") // Unrecognised COFF file header machine value of 0x6170.
	file, err := pe.Open("hello")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", file)

	// {
	// 	Machine:34404
	// 	NumberOfSections:12
	// 	TimeDateStamp:0
	// 	PointerToSymbolTable:2193408
	// 	NumberOfSymbols:3607
	// 	SizeOfOptionalHeader:240
	// 	Characteristics:547
	// }
	log.Printf("%+v\n", file.FileHeader)
}
