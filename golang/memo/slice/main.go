package main

import (
	"fmt"
	"log"
)

func main() {

	// 将 fmt 全部注释，输出 bb; 取消注释，输出 ab，且 s 的长度和容量均是 0 ，为什么呢？
	// https://gocn.io/question/1852
	// []byte("")的cap到底应该是多少，属于未定义行为，允许编译器根据上下文做优化和调整
	// s := []byte("")

	s := make([]byte, 0, 32) // 指定容量后，输出 bb
	fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)

	s1 := append(s, 'a')
	fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)
	fmt.Printf("%d, %d, %p\n", len(s1), cap(s1), s1)

	s2 := append(s, 'b')
	fmt.Printf("%d, %d, %p\n", len(s), cap(s), s)
	fmt.Printf("%d, %d, %p\n", len(s2), cap(s2), s2)

	log.Print(string(s1), string(s2))

}
