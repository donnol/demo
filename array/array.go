package main

import (
	"fmt"
)

func main() {
	var array = [8]int{}
	for i := range array {
		array[i] = i
	}
	format(&array, &array[0], len(array), cap(array), array)

	var slice = array[:2]
	format(&slice, &slice[0], len(slice), cap(slice), slice)

	// mod
	slice[0] = 10
	format(&array, &array[0], len(array), cap(array), array)
	format(&slice, &slice[0], len(slice), cap(slice), slice)

	// add one
	// slice[2] = 12 // panic: runtime error: index out of range
	slice = append(slice, 12)

	format(&array, &array[0], len(array), cap(array), array)
	format(&slice, &slice[0], len(slice), cap(slice), slice)

	// add more
	slice = append(slice, []int{13, 14, 15, 16, 17}...)
	format(&array, &array[0], len(array), cap(array), array)
	format(&slice, &slice[0], len(slice), cap(slice), slice)

	// over and expansion
	slice = append(slice, 18)
	format(&array, &array[0], len(array), cap(array), array)
	format(&slice, &slice[0], len(slice), cap(slice), slice)

	// mod
	slice[0] = 20
	format(&array, &array[0], len(array), cap(array), array)
	format(&slice, &slice[0], len(slice), cap(slice), slice)

	// non-constant array bound l
	// var l = 8
	// array1 := [l]int{}
	// for i := 0; i < l; i++ {
	// 	array1[i] = i
	// }
}

func format(args ...interface{}) {
	fmt.Printf("ptr is: %p, first ele ptr is: %p, len is: %d, cap is: %d, val is: %v\n", args...)
}
