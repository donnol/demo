package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Println("你好")
}

var Greeter greeting
