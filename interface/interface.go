package main

import (
	"fmt"
)

func main() {
	i := NewI()
	i.Run()                  // 接口声明的方法，可以直接调用
	i.(*M).Stop()            // 调用接口没有但是结构体有的方法，需要断言
	fmt.Println(i.(*M).Name) // 使用结构体的字段，也需要断言
}

type I interface {
	Run() error
}

type M struct {
	Name string
}

func (m *M) Run() error {
	fmt.Println(m.Name, "is running.")
	return nil
}

func (m *M) Stop() error {
	fmt.Println(m.Name, "is stop.")
	return nil
}

func NewI() I {
	return &M{"jd"}
}
