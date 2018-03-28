// 生成汇编文件
// 	1. go build -gcflags '-S' main.go > main.s 2>&1 构建程序的同时生成
// 	2. go tool compile -S main.go > main.s 源码生成，同时还会生成 main.o 文件
//  3. go tool objdump -S main > main.s 反编译二进制可执行文件生成，结果会有点不同
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world.")
}
