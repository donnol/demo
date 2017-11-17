// 接口的私有方法是否需要继承
package main

import "fmt"

func main() {
	m := M{}
	Print(m)
}

type Runner interface {
	Run() error
	run() error
}

type M struct {
}

func (m M) run() error {
	fmt.Println("m is runing.")
	return nil
}

func (m M) Run() error {
	return m.run()
}

func Print(r Runner) {
	r.Run()
}
