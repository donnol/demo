package main

import "fmt"

func main() {
	b := B{A{name: "jd"}, "jj"}
	h := b.Hello()
	fmt.Println(h)
}

type A struct {
	name string
}

func (a A) Hello() string {
	return fmt.Sprintln("a: my name is", a.name)
}

type B struct {
	A
	name string
}

func (b B) Hello() string {
	return fmt.Sprintln("b: my name is", b.name)
}
