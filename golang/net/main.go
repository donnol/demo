package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取本机地址信息
	inters, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", inters)

	for _, inter := range inters {
		fmt.Printf("%s, %s\n", inter.Network(), inter.String())
		// 文档里说返回值可以直接用来Dial，其实还要看具体实现，这里是不行的
		// conn, err := net.Dial(inter.Network(), inter.String())
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Printf("conn: %#v\n", conn)
	}
}
