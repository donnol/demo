package main

import (
	"fmt"
	"log"
)

func main() {

	// 将 fmt 全部注释，输出 bb; 取消注释，输出 ab，且 s 的长度和容量均是 0 ，为什么呢？
	//	逃逸分析
	// 	s 虽然是在 main 里声明及使用，但当它传入 fmt.Printf 时，会转型为 interface{}，而 interface{} 类型中包含指向数据的地址，导致了内存重分配，所以 s 最后逃逸到堆里
	//	可使用 go build -gcflags '-m -m -l' main.go 查看具体情况

	// https://gocn.io/question/1852
	// 后续解析：https://www.do1618.com/archives/1328/go-%E5%86%85%E5%AD%98%E9%80%83%E9%80%B8%E8%AF%A6%E7%BB%86%E5%88%86%E6%9E%90/
	// s := []byte("") // 这里是将一个空字符串转型为字节切片，转型时调用了底层的 stringtoslicebyte，这个函数规定：如果字符串的大小没有超过 32 长度的大小，则默认分配一个 32 长度的 buf

	// s := make([]byte, 0, 32) // 指定容量后，输出 bb
	s := []byte{} // 零值，切片的长度和容量都是 0，输出 ab
	fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)

	s1 := append(s, 'a')
	fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)
	fmt.Printf("%d, %d, %p\n", len(s1), cap(s1), s1)

	s2 := append(s, 'b')
	fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)
	fmt.Printf("%d, %d, %p\n", len(s2), cap(s2), s2)

	log.Print(string(s1), string(s2))

}
