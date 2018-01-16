// Go 里面的函数传参都是值
package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	fmt.Println("========================")
	var i = 1
	var f = 0.1
	var s = "hello"

	simple(i, f, s)
	fmt.Println("value is", i, f, s)
	fmt.Println("addr is", &i, &f, &s)

	fmt.Println("========================")
	a := [2]int{0, 1}
	slice := []int{0, 1, 2}
	array(a, slice)
	fmt.Println("value is", a, slice)
	fmt.Printf("addr is  is %p, %p\n", &a, &slice)

	fmt.Println("========================")
	m := map[int]int{0: 0}
	passMap(m)
	fmt.Println("value is", m)
	fmt.Printf("addr is %p\n", &m)

	fmt.Println("========================")
	var c = make(chan int)
	go func() {
		passChan(c)
	}()
	if c != nil {
		fmt.Println("value is", <-c) // 这里还是不变，所以会一直阻塞，最后出现死锁错误
	}
	fmt.Printf("addr is %p\n", &c)

	fmt.Println("========================")
	function := func() {
		fmt.Println("i am a function")
	}
	passFunc(function)
	fmt.Printf("value is %+v\n", function)
	fmt.Printf("addr is %p\n", &function)

	fmt.Println("========================")
	var buf = make([]byte, 12)
	buffer := bytes.NewBuffer(buf)
	passInterface(buffer)
	fmt.Printf("value is '%+v'\n", buffer)
	fmt.Printf("addr is %p\n", &buffer)

	fmt.Println("========================")
	passEmptyInterface(s)
	fmt.Println("value is", s)
	fmt.Printf("addr is %p\n", &s)
}

// 简单类型
func simple(i int, f float64, s string) {
	i = 2
	f = 0.2
	s = "world"
	fmt.Println("func value is", i, f, s)
	fmt.Println("func addr is", &i, &f, &s)
}

// 数组 切片
func array(a [2]int, slice []int) {
	a[0] = 1
	// slice[0] = 1 // 两个 slice 共用了一个数组
	slice = append(slice, 3) // 切片扩容时新建了数组
	fmt.Println("func value is", a, slice)
	fmt.Printf("func addr is %p, %p\n", &a, &slice)
}

// map
func passMap(m map[int]int) {
	// m[0] = 1 // 两个 map 共用底层数据，外面的 map 也会更改
	m = nil // 当把 m 设为 nil 后，外面的 map 保持不变
	fmt.Println("func value is", m)
	fmt.Printf("func addr is %p\n", &m)
}

// chan
func passChan(c chan int) {
	// c <- 1 // 两个 chan 共用底层数据，这里写的 外面的 chan 一样可以接收
	// c = nil // 当把 c 设为 nil 后，外面的 c 仍然不变
	close(c) // 当把 c close 后，外面的 c 会关闭，但是依然会接收到 int 类型零值 0
	fmt.Printf("func addr is %p\n", &c)
}

// func
func passFunc(f func()) {
	// f() // 两个 func 共用底层数据
	f = nil // 当把 f 设为 nil，外面的 func 仍然不变
	fmt.Printf("func value is %+v\n", f)
	fmt.Printf("func addr is %p\n", &f)
}

// interface
func passInterface(w io.Writer) {
	// w.Write([]byte("hello"))
	w = nil
	fmt.Printf("func value is '%+v'\n", w)
	fmt.Printf("func addr is %p\n", &w)
}

// empty interface
func passEmptyInterface(e interface{}) {
	e = nil
	fmt.Println("func value is", e)
	fmt.Println("func addr is", &e)
}
