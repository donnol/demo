package main

import (
	"hash/crc32"
	"log"
)

func main() {
	data := []byte("Hello, I am jd")
	tab := crc32.IEEETable
	sum1 := crc32.Checksum(data, tab)
	sum2 := crc32.ChecksumIEEE(data)
	log.Printf("%d, %d\n", sum1, sum2)

	tab = crc32.MakeTable(0x123456)
	// log.Printf("%+v\n%+v\n", tab, crc32.IEEETable)

	sum1 = crc32.Checksum(data, tab)
	sum2 = crc32.ChecksumIEEE(data)
	log.Printf("%d, %d\n", sum1, sum2)
}
