package main

import (
	"bytes"
	"io"
	"log"
)

func main() {
	src := []byte("Hello, I am jd")
	srcReader := bytes.NewBuffer(src)
	dst := []byte{}
	dstWriter := bytes.NewBuffer(dst)

	// io.Copy
	// n, err := io.Copy(dstWriter, srcReader) // srcReader 的内容被读走，并写到 dstWriter
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(n, len(src), len(dst), src, dst, srcReader.String(), dstWriter.String())

	// io.CopyBuffer
	// buf := make([]byte, len(src))
	// n, err := io.CopyBuffer(dstWriter, srcReader, buf) // 通过 buf，而不是自己分配一个
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(n, len(src), len(dst), src, dst, srcReader.String(), dstWriter.String(), buf)

	// io.CopyN
	n, err := io.CopyN(dstWriter, srcReader, 8) // 读走 8 个字节
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, len(src), len(dst), src, dst, srcReader.String(), dstWriter.String())

}
