package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := func() int {
		return 1
	}
	b := func() int {
		return 1
	}
	c := b

	r := reflect.DeepEqual(a, b)
	fmt.Println(r)

	r = reflect.DeepEqual(c, b)
	fmt.Println(r)

	var d func() int
	var e func() int
	r = reflect.DeepEqual(d, e)
	fmt.Println(r)
}
