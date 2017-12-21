package main

import (
	"log"

	"alias2"
	// "./alias"
)

// A A 原始类型
type A string

// B B 类型别名
// type B = A

// 子目录下的可以找到
// type m = alias.M

// vendor 目录的无法找到
// type m2 = alias2.M

func main() {
	var a A = "a"

	log.Printf("%s\n", a)

	// var b B = "b"

	// log.Printf("%s\n", b)

	// var mval m = "m"
	// log.Printf("%s\n", mval)

	// var mval2 m2 = "m2"
	// log.Printf("%s\n", mval2)

	// vendor 目录 直接使用 还是无法找到
	var m2 alias2.M = "m2"
	log.Printf("%s\n", m2)
}
