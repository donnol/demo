#!/usr/bin/env gorun

package main

import (
	"fmt"
	"unsafe"
)

func float64bits(f float64) uint64 {
	p := &f                 // 变量f地址
	up := unsafe.Pointer(p) // 转型
	upi := (*uint64)(up)    // 转10进制
	// fmt.Printf("%p, %p, %d, %d, %d\n", p, up, upi, *upi, *up) // invalid indirect of up (type unsafe.Pointer)
	fmt.Printf("%p, %p, %d, %d\n", p, up, upi, *upi)
	return *upi
}

func main() {
	for _, v := range []float64{
		3.14,
		2.13,
		0.01,
	} {
		float64bits(v)
		fmt.Printf("%d\n", unsafe.Alignof(v))
		var s = struct {
			f1 float64
			f2 float64
		}{
			v,
			v,
		}
		fmt.Printf("%d\n", unsafe.Offsetof(s.f2))
		fmt.Printf("%d\n", unsafe.Sizeof(v))
	}
}
