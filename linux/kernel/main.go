package main

import (
	"fmt"
)

var a = 1

func init() {
	fmt.Printf("a addr: %p\n", &a)
}

func main() {
	var b = 1
	fmt.Printf("b addr: %p\n", &b)
}
