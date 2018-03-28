package main

import (
	"fmt"
)

func main() {
	m := make(map[int]int)
	m[0] = 0
	f(m)
	fmt.Println(m)
}

func f(m map[int]int) {
	m[1] = 1
}
