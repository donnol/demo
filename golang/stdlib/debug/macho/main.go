package main

import (
	"debug/macho"
	"log"
)

func main() {
	file, err := macho.Open("hello") // invalid magic number in record at byte 0x0
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", file)

	// {
	// 	Magic:4277009103
	// 	Cpu:CpuAmd64 // 5 种
	// 	SubCpu:3
	// 	Type:Exec
	// 	Ncmd:10
	// 	Cmdsz:2296
	// 	Flags:1
	// }
	log.Printf("%+v\n", file.FileHeader)
}
