package main

import (
	"fmt"
)

func main() {
	a := []int{4, 3, 6, 5, 1, 2, 9, 7, 8, 10}
	r := insertSort(a)
	fmt.Println(r)
}

// 插入排序：假设前面k个元素已经有序，第k+1个元素应该插入到哪个位置呢？
func insertSort(a []int) []int {
	alen := len(a)
	for j := 1; j < alen; j++ { // 从第二个元素开始，到最后一个元素
		key := a[j]                // 第二个元素
		i := j - 1                 // 前面一个元素的下标
		for i >= 0 && a[i] > key { // 如果前一个元素值大于后一个元素值
			a[i+1] = a[i] // 就将大元素值移到后面
			i = i - 1     // 下标减一，也就是比 j 小的小标组成的数组都要和 j 下标的值比较
		}
		a[i+1] = key // 如果没有执行上面的循环，还是原来的位置，原来的值；
	}
	return a
}

// 4, 3, 6, 5, 1, 2, 9, 7, 8, 10 -- 初始数组
// 3, 4, 6, 5, 1, 2, 9, 7, 8, 10 -- j = 1
// 3, 4, 6, 5, 1, 2, 9, 7, 8, 10 -- j = 2
// 3, 4, 6, 6, 1, 2, 9, 7, 8, 10 -- j = 3, i = 2
// 3, 4, 6, 6, 1, 2, 9, 7, 8, 10 -- j = 3, i = 1
// 3, 4, 5, 6, 1, 2, 9, 7, 8, 10 -- j = 3, i = 1
// 3, 4, 5, 6, 6, 2, 9, 7, 8, 10 -- j = 4, i = 3
// 3, 4, 5, 5, 6, 2, 9, 7, 8, 10 -- j = 4, i = 2
// 3, 4, 4, 5, 6, 2, 9, 7, 8, 10 -- j = 4, i = 1
// 1, 3, 4, 5, 6, 2, 9, 7, 8, 10 -- j = 4, i = 0
