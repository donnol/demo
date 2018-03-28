package main

import (
	"fmt"
)

func main() {
	s := "hello"
	fmt.Println(s)

	escape0()
	escape1()
	escape2()
}

type S struct{}

func escape0() {
	var x S
	y := &x
	_ = *identity(y)
}

func identity(z *S) *S {
	return z
}

type S1 struct{}

func escape1() {
	var x S
	_ = *ref(x)
}

func ref(z S) *S {
	return &z
}

type S2 struct {
	M *int
}

func escape2() {
	var i int
	refStruct(i)
}

func refStruct(y int) (z S2) {
	z.M = &y
	return z
}

func escape3() {
	var i int
	refStruct3(&i)
}

func refStruct3(i *int) (z S2) {
	z.M = i
	return z
}

func escape4() {
	var x S2
	var i int
	ref3(&i, &x)
}

func ref3(y *int, z *S2) {
	z.M = y
}
