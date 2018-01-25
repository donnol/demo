package main

import "fmt"

func main() {
	// 如何给一个bool指针赋值 ?
	var bp *bool

	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x487c6f]
	// 因为bp和sp的值是nil，指向地址0x0，所以在使用*bp和*sp时都会报非法内存的错误
	// 所以，在使用指针变量时，最好先用new来初始化
	// fmt.Println(*bp)

	// 同上
	// *bp = true
	// fmt.Println(*bp)

	// 同上
	// var b1 = true
	// *bp = b1
	// fmt.Println(*bp)

	// ok
	var b2 = true
	bp = &b2
	fmt.Println(*bp)

	// ok
	var b3 = true
	var bp1 = &b3
	*bp1 = false
	fmt.Println(*bp1)

	// ok
	var bp2 = new(bool)
	fmt.Println(*bp2)

	// 切片呢？
	var sp *[]int

	// 同上
	// fmt.Println(*sp)

	// *sp = []int{1, 2, 3}
	// fmt.Println(*sp)

	// var s1 = []int{1, 2, 3}
	// *sp = s1
	// fmt.Println(*sp)

	var s1 = []int{1, 2, 3}
	sp = &s1
	fmt.Println(*sp)

	var sp1 = new([]int)
	fmt.Println(*sp1)
}
