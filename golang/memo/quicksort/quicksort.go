package main

import (
	"fmt"
)

func main() {
	a := []int{3, 2, 6, 5, 9, 8, 1, 7, 4}
	quicksort(a, 0, 8) // 传入数组，第一个元素下标，最后一个元素下标
	fmt.Println(a)
}

// func quicksort(a []int, p, r int) {
// 	if p < r {
// 		q := partition(a, p, r)
// 		quicksort(a, p, q)
// 		quicksort(a, q, r)
// 	}
// }

// func partition(a []int, p, r int) int {
// 	x := a[r-1]
// 	i := p - 1
// 	for j := 0; j < r-1; j++ {
// 		if a[j] <= x {
// 			i++
// 			a[i] = x
// 		}
// 	}
// 	// a[i] = a[i-1]
// 	return i + 1
// }

// 来自算法导论
// QuickSort(A,p,r)
// 	if p < r
// 		q = Partition(A, p, r)
// 		QuickSort(A, p, q-1)
// 		QuickSort(Q, q+1, r)
func quicksort(s []int, l, r int) {
	if r > l {
		p := partition(s, l, r) // 基准的位置确定了
		quicksort(s, l, p-1)    // 基准左边的元素继续排序
		quicksort(s, p+1, r)    // 右边的也一样
	}
}

// 原地交换
// Partition(A, p, r)
// 	x = A[r]
// 	i = p-1
// 	for  j = p  to  r-1
// 		if  A[j] <= x
// 			then  i = i+1
// 				exchange A[i] <-> A[j]
// 	exchange A[i+1]<->A[r]
// 	return i+1
func partition(s []int, l, r int) int {
	pivot := s[r]               // 最后的元素作为基准
	i := l - 1                  // 小于基准的元素的 索引
	for j := l; j <= r-1; j++ { // 从左到右遍历 除基准外的元素
		if s[j] <= pivot { // 小于基准的元素 -- 决定什么时候交换
			i++                     // 决定交换的元素下标
			s[j], s[i] = s[i], s[j] // 将元素交换到左边的位置
		}
	}
	s[i+1], s[r] = s[r], s[i+1] // 将基准放到合适的位置 -- 下标 i+1 前面的元素值都是比 s[r] 小的
	return i + 1
}
