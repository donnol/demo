package main

import "fmt"

type f func(int, int)

type I interface {
	// sub(int, int)
	// sub(a int, b int)
}

type M struct {
}

func (m M) sub(a int, b int) {
	fmt.Println(a + b)
}

func main() {
	// if 'abc' < 'ABC' {
	if 'a' < 'A' {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}

	// if 8 > 4 > 2 {
	// 	fmt.Printf("true\n")
	// } else {
	// 	fmt.Printf("false\n")
	// }

	if 9 < 1 && 10 > 9 || 2 > 1 {
		fmt.Printf("true\n")
	} else {
		fmt.Printf("false\n")
	}
}
