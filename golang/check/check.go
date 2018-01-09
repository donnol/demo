package main

import (
	"fmt"
	"strings"
)

func main() {
	check(0)

	a := newArray(10)
	printPtrVal(a)

	check(1)
	printPtrVal(a)

	check(2)
	printPtrVal(a)
}

var checkNum int

func check(i int) {
	checkNum = i
}

func newArray(n int) []*int {
	r := []*int{}
	for i := 0; i < n; i++ {
		r = append(r, &checkNum)
	}
	return r
}

func printPtrVal(a []*int) {
	str := "["
	for _, single := range a {
		str += fmt.Sprintf("%d ", *single)
	}
	str = strings.TrimRight(str, " ")
	str += "]"
	fmt.Println(str)
}
