package main

import (
	"container/heap"
	"fmt"
)

func main() {
	heapTest()
}

// Heap 整型堆
type Heap struct {
	values []int
}

func (h Heap) Swap(i, j int) {
	if h.Len() <= 1 {
		return
	}
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h Heap) Len() int {
	return len(h.values)
}

func (h Heap) Less(i, j int) bool {
	if h.Len() <= 1 {
		return true
	}
	return h.values[i] < h.values[j]
}

// Push 推入
func (h *Heap) Push(v interface{}) {
	h.values = append(h.values, v.(int))
}

// Pop 弹出
func (h *Heap) Pop() interface{} {
	if h.Len() == 0 {
		panic("heap is empty")
	}
	i := h.Len() - 1
	pop := h.values[i]
	h.values = h.values[:i]
	return pop
}

func heapTest() {
	h := &Heap{values: []int{3, 2, 1}}
	heap.Init(h) // 将小的放前面
	fmt.Println("init:", h)

	v := heap.Pop(h) // 将前面的值弹出来
	fmt.Println("pop:", h, v)
	v = 4
	heap.Push(h, v)
	fmt.Println("push:", h, v)
	v = heap.Pop(h)
	fmt.Println("pop:", h, v)
	v = heap.Pop(h)
	fmt.Println("pop:", h, v)
	v = heap.Pop(h)
	fmt.Println("pop:", h, v)
	// v = heap.Pop(h)
	// fmt.Println("pop:", h, v)
}
