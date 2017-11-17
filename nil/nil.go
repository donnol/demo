// nil is 0x0
package main

import (
	"fmt"
)

func main() {
	// fmt.Println(nil == nil)

	var m *map[int]int
	var s *[]int
	fmt.Printf("%p, %p\n", m, s)

	// fmt.Println(m == s)

	var a *int = nil
	fmt.Printf("%p\n", a)
}
