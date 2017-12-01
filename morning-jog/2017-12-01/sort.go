package p20171201

// Sort 排序
func Sort(s []int) []int {
	l, r := 0, len(s)-1
	quicksort(s, l, r)
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
		p := partition(s, l, r)
		quicksort(s, l, p-1)
		quicksort(s, p+1, r)
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
	pivot := s[r]
	i := l - 1
	for j := l; j <= r-1; j++ {
		if s[j] <= pivot {
			i++
			s[j], s[i] = s[i], s[j]
		}
	}
	s[i+1], s[r] = s[r], s[i+1]
	return i + 1
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
