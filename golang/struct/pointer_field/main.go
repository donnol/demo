package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	log.Printf("%s\n", (&bytes.Buffer{}).Bytes()) // bytes.Buffer 里不存在结构体指针字段

	log.Printf("%s\n", (&m{}).String())
}

type m struct {
	p *struct { // 这个就是结构体里的结构体指针字段
		q int
	}
}

func (m *m) String() string {
	fmt.Printf("%v\n", m.p) // <nil>

	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x0 pc=0x498448]
	return fmt.Sprintf("%d\n", m.p.q) // m.p默认初始化为nil，所以打印m.p.q时出错
}
