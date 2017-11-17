package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 6, 5, 9, 8, 1, 7, 4}
	quicksort(a, 0, 8)
	fmt.Println(a)
}

func quicksort(a []int, p, r int) {
	if p < r {
		q := partition(a, p, r)
		quicksort(a, p, q)
		quicksort(a, q, r)
	}
}

func partition(a []int, p, r int) int {
	x := a[r-1]
	i := p - 1
	for j := 0; j < r-1; j++ {
		if a[j] <= x {
			i++
			a[i] = x
		}
	}
	// a[i] = a[i-1]
	return i + 1
}
