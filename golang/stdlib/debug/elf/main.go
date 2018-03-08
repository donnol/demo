package main

import (
	"debug/elf"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// file, err := elf.Open("main.go") // bad magic number '[112 97 99 107]' in record at byte 0x0
	file, err := elf.Open("hello")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err = file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	log.Println(file)

	data, err := file.DWARF()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", data)

	dynSym, err := file.DynamicSymbols()
	if err != nil {
		if err != elf.ErrNoSymbols {
			log.Fatal(err)
		}
	}
	log.Printf("%+v\n", dynSym)

	sym, err := file.Symbols()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", sym)

	// {
	// Class:ELFCLASS64 // 另一种是 ELFCLASS32
	// Data:ELFDATA2LSB // 小端(最低有效位（Least Significant Byte）占用最低地址)
	//  另一个是 ELFDATA2MSB 大端(最高有效位（Most Significant Byte）占用最低地址)
	// Version:EV_CURRENT
	// OSABI:ELFOSABI_NONE // 总共 16 种
	// ABIVersion:0
	// ByteOrder:LittleEndian
	// Type:ET_EXEC // 8 种
	// Machine:EM_X86_64 // 机器架构 非常多
	// Entry:4527520
	// }
	log.Printf("%+v\n", file.FileHeader)
}
