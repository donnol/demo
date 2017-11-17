package main

import (
	"fmt"
)

func main() {
	var m = M{P{}}
	m.Run()
}

type I interface {
	Run()
}

type M struct {
	I
}

type P struct {
}

func (p P) Run() {
	fmt.Println("p is running.")
}
