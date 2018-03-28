// 位运算
package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		// fmt.Printf("有符号正数：%04b, %04b\n", int(i), ^int(i))
		fmt.Printf("无符号正数：%04b, %04b\n", uint(i), ^uint(i))
		// fmt.Printf("有符号负数：%04b, %04b\n", int(-i), ^int(-i))
		fmt.Printf("无符号负数：%04b, %04b\n", uint(-i), ^uint(-i))
	}
}
