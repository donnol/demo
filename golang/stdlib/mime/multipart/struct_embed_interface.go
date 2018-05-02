package main

import (
	"io"
	"log"
	"mime/multipart"
)

func structEmbedInterfaceMeaning() {
	// var _ io.Reader = multipart.File // 如何确定一个接口内嵌了另一个接口呢？
	var _ io.Reader = &M{}
	var _ io.ReaderAt = M{}
	var _ io.Seeker = M{}
	var _ io.Closer = M{}

	m := &M{}
	p := make([]byte, 0, 8)
	n, err := m.Read(p) // 因为 1，所以m实例有Read方法，但是，会报错: panic: runtime error: invalid memory address or nil pointer dereference
	// [signal 0xc0000005 code=0x0 addr=0x20 pc=0x4ba4b7]
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
}

// M M 类型
type M struct {
	// 结构体内嵌接口主要用于：
	// 	当 M 只想重载 File 接口的其中一个方法时；
	// 	或者 M 需要以 File 接口参数被传入另一个函数中
	// 参考：https://stackoverflow.com/questions/24537443/meaning-of-a-struct-with-embedded-anonymous-interface
	multipart.File // 1 嵌入了 File 接口
}

func (m *M) Read(p []byte) (n int, err error) {
	return
}
