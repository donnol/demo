package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 6, 5, 1, 2, 9, 7, 8, 10}
	r := insertSort(a)
	fmt.Println(r)
}

func insertSort(a []int) []int {
	alen := len(a)
	for j := 1; j < alen; j++ {
		key := a[j]
		i := j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}
	return a
}
