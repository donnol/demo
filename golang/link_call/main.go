package main

import (
	"fmt"
)

type model struct {
	name string
	age  int
}

func (m *model) SetName(name string) *model {
	m.name = name
	return m
}

func (m *model) SetAge(age int) *model {
	m.age = age
	return m
}

func (m *model) String() string {
	return fmt.Sprintf("name is %s, age is %d\n", m.name, m.age)
}

func main() {
	m1 := &model{}
	fmt.Println(m1.SetName("jd"))
	m2 := &model{}
	fmt.Println(m2.SetAge(20))
	m3 := &model{}
	fmt.Println(m3.SetName("jd").SetAge(20))
}
