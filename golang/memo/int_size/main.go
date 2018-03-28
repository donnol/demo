// int 和 uint 的长度，4字节，还是8字节呢？
package main

import (
	"log"
	"math"
	"runtime"
	"strconv"
)

func main() {
	// var a int = math.MaxUint64 // constant 18446744073709551615 overflows int

	var a uint = math.MaxUint64
	log.Printf("max uint is: %d\n", a) // 18446744073709551615

	log.Printf("int size is: %s, %d\n", runtime.GOARCH, strconv.IntSize)

	x := ^uint32(0)
	i := int(x) // 如果在 32 位系统里，i 的值是 -1
	log.Printf("0x%x, %X", x, i)
}
