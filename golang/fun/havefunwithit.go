package main

import "fmt"

func main() {
	type A *[]*A
	a := &[]A{[]A{{}}[0]}

	fmt.Println(*a)
}
