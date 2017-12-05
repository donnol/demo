package p20171201

import (
	"sync"
)

// Sort 排序
// 排序过程可视化 https://www.cs.usfca.edu/~galles/visualization/ComparisonSort.html
func Sort(s []int) []int {
	l, r := 0, len(s)-1
	quicksort(s, l, r)
	return s
}

// SortConcurrent 并行排序
func SortConcurrent(s []int) []int {
	l, r := 0, len(s)-1
	var wg sync.WaitGroup
	wg.Add(1)
	go quicksortConcurrent(&wg, s, l, r)
	wg.Wait()
	return s
}

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
	i := l - 1                  // 小于基准的元素索引
	for j := l; j <= r-1; j++ { // 从左到右遍历 除基准外的元素
		if s[j] <= pivot { // 小于基准的元素
			i++
			s[j], s[i] = s[i], s[j] // 将元素交换到左边的位置
		}
	}
	s[i+1], s[r] = s[r], s[i+1] // 将基准放到合适的位置
	return i + 1
}

// 并行计算
func quicksortConcurrent(wg *sync.WaitGroup, s []int, l, r int) {
	defer wg.Done()
	if r > l {
		p := partition(s, l, r)
		wg.Add(2)
		go quicksortConcurrent(wg, s, l, p-1)
		go quicksortConcurrent(wg, s, p+1, r)
	}
}

// wiki https://zh.wikipedia.org/wiki/%E5%BF%AB%E9%80%9F%E6%8E%92%E5%BA%8F
func qsort(data []int) {
	if len(data) <= 1 {
		return
	}
	mid, i := data[0], 1
	head, tail := 0, len(data)-1
	for i = 1; i <= tail; {
		if data[i] > mid {
			data[i], data[tail] = data[tail], data[i]
			tail--
		} else {
			data[i], data[head] = data[head], data[i]
			head++
			i++
		}
	}
	data[head] = mid
	qsort(data[:head])
	qsort(data[head+1:])
}
