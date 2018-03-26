package main

import (
	"go/types"
)

type t struct{}

func (t t) Underlying() types.Type {
	return t
}

func (t t) String() string {
	return "t"
}

func main() {
	var _ types.Type = t{}
}
