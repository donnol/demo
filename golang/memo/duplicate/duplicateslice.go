package main

import (
	"fmt"
)

func main() {
	r := makeDuplicateSlice(10, 15)
	fmt.Println(r)
}

func makeDuplicateSlice(len, cap int) *[]*int {
	r := make([]*int, len, cap)
	for i := 0; i < len; i++ {
		r[i] = &i
	}
	fmt.Println(r)
	return &r
}
