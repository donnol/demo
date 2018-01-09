package main

import (
	"fmt"
	"reflect"
)

func main() {
	type MyInt int

	var x MyInt = 1
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Println(t, v.Kind())

	vi := v.Interface().(MyInt)
	fmt.Println(vi)

	var y float64 = 3.4
	yv := reflect.ValueOf(y)
	fmt.Println(yv.CanSet())

	var yp *float64 = &y
	ypv := reflect.ValueOf(yp).Elem()
	fmt.Println(ypv.CanSet())
	ypv.SetFloat(5.5)
	fmt.Println(y, *yp)
}
