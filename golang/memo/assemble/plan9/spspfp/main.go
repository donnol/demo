package main

import (
	"fmt"
)

func output(int) (int, int, int) // 汇编函数声明 -- 这个注释不是必须的

func main() {
	a, b, c := output(987654321)
	fmt.Println(a, b, c)
}
