package main

import (
	"a"
	"fmt"
)

var s a.A = "a"

// 只要引入了 类型别名 的变量，就会导致该包的变量都 godef 无法找到定义，使用 gogetdoc 可以
var s1 a.B = "b"

func main() {
	fmt.Println(s)
	fmt.Println(s1)
}
