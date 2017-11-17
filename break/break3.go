package main

import (
	"fmt"
)

func main() {
	fmt.Println("[")
	// RET: // invalid break label RET
	temp := 7
RET:
	for i := 0; i < temp; i++ {
		fmt.Println(i)
		if i == 3 {
			break RET
		}
	}
	fmt.Println("]")
}
