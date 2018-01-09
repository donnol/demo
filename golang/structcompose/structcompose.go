package main

import "fmt"

func main() {
	c := Compose{S{"jd"}, "lau"}
	fmt.Println(c)
	c.Print()
	c.S.Print()
}

type S struct {
	Name string
}

func (s S) Print() {
	fmt.Println(s.Name)
}

type Compose struct {
	S
	Name string
}

func (c Compose) Print() {
	fmt.Println(c.Name)
}
