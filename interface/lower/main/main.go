package main

import (
	"demo/interface/lower"
	"fmt"
)

func main() {
	var _ lower.i = &M{"jd"}
	// 	src\demo\interface\lower\main\main.go:9: cannot refer to unexported name lower.i
	// src\demo\interface\lower\main\main.go:9: undefined: lower.i
}

type M struct {
	Name string
}

func (m *M) Run() error {
	fmt.Println(m.Name, "is ruuning.")
	return nil
}
