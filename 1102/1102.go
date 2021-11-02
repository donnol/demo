package main

import "fmt"

func main() {
	indexes := getIndexes([]int{7, 1, 5, 3, 6, 4, 3}, 3)
	fmt.Println(indexes)
}

// arr = [7,1,5,3,6,4,3], target = 3
func getIndexes(arr []int, target int) []int {
	r := make([]int, 0, len(arr))
	for i, s := range arr {
		if s != target {
			continue
		}
		r = append(r, i)
	}
	return r
}
