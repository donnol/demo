package main

import (
	"debug/dwarf"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	abbrev, aranges, frame, info, line, pubnames, ranges, str := []byte("hello"), []byte("I"), []byte("am"), []byte("jd"), []byte("what"), []byte("is"), []byte("your"), []byte("name")
	// Rather than calling this function directly, clients should typically use the DWARF method of the File type of the appropriate package debug/elf, debug/macho, or debug/pe.
	// 不要直接调用这个方法，这个方法是提供给其它包调用的
	data, err := dwarf.New(abbrev, aranges, frame, info, line, pubnames, ranges, str)
	if err != nil {
		log.Fatal(err) // decoding dwarf section info at offset 0x2: too short
	}
	log.Println(data)
}
