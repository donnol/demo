package main

import "fmt"

func add(x, y int) int { // 汇编调用了这里的add函数
	return x + y
}

func output(a, b int) int

func main() {
	s := output(10, 13)
	fmt.Println(s)
}
