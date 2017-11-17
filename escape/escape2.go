package main

import (
	"fmt"
)

func main() {
	m := NewM("jd")
	fmt.Println(m)

	m2 := NewM2("jd")
	fmt.Println(m2)
}

type M struct {
	Name string
}

func NewM(name string) *M {
	newName := name
	r := &M{
		Name: newName,
	}
	return r
}

func NewM2(name string) M {
	newName := name
	r := M{
		Name: newName,
	}
	return r
}
