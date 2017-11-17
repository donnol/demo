package main

import "fmt"

func main() {
loop:
	for i := 0; i < 10; i++ {
		// loop:
		for {
			// if i%2 == 0 {
			fmt.Println("in : ", i)
			// }
			// goto loop
			break loop
			// break

			// fmt.Println("in after break : ", i)
		}

		fmt.Println("out: ", i)
	}
}
