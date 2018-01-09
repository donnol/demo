package main

import (
	"fmt"
)

func main() {
	var a M = A{"a"}
	var b M = B{A{"a"}, "b"}
	fmt.Println(a.String())
	fmt.Println(b.String())
}

type M interface {
	String() string
}

type A struct {
	Name string
}

func (a A) String() string {
	return a.Name
}

type B struct {
	A
	Name string
}
