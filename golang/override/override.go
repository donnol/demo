package main

import (
	"fmt"
)

type M struct {
	Name string
}

func (m *M) String() string {
	return fmt.Sprintf("m.name is %s", m.Name)
}

func main() {
	m := &M{"jd"}
	m.String = func() string {
		return "hi"
	}
	// cannot assign to m.String

	m.String()
}
