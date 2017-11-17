package main

import "fmt"

func main() {
	var c = 1 + b
	var b int
	fmt.Println(b, c)
}

var b int

func init() {
	b = 2
}
