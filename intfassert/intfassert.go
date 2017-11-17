package main

import "fmt"

func main() {
	// Intf().(I).Run()
	// Intf().(M).Run()
	Intf().Run() // Intf().Run undefined (type interface {} is interface with no methods) // 空接口，没有任何方法
}

func Intf() interface{} {
	m := M{"jd"}
	return m
}

type I interface {
	Run() error
}

type M struct {
	Name string
}

func (m M) Run() error {
	fmt.Println(m.Name, "is running...")
	return nil
}
