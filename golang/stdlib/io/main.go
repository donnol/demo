package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

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

	// pipe
	pr, pw := io.Pipe() // 往 pw 写东西时，pr 也会读到
	defer pr.Close()
	defer pw.Close()
	go func() {
		pw.Write(src)
	}()
	prData := make([]byte, len(src))
	_, err = pr.Read(prData)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(src), string(prData))

	// limitReader
	buf := bytes.NewBuffer(src)
	lr := io.LimitedReader{ // 最多只能读到 R 的前面 N 个，已经读够后，报错 EOF
		R: buf, // 长度 > 8
		N: 8,   // 只能读到 8 个
	}
	lrData := make([]byte, 7) // 第一次读了 7 个
	_, err = lr.Read(lrData)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("'%s', '%s'", string(src), string(lrData))
	lrData = make([]byte, 8)
	_, err = lr.Read(lrData) // 还剩 1 个
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("'%s', '%s'", string(src), string(lrData))

	// multiReader -- 级联
	a := strings.NewReader("a")
	b := strings.NewReader("b")
	c := strings.NewReader("c")
	z := io.MultiReader(a, b, c)
	// zData := make([]byte, a.Len()+b.Len()+c.Len())
	// log.Printf("%d, %d, %d, %d, %d\n", a.Len(), b.Len(), c.Len(), len(zData), cap(zData))
	// _, err = z.Read(zData)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("'%s'", zData) // 'a  '
	if _, err = io.Copy(os.Stdout, z); err != nil { // abc
		log.Fatal(err)
	}
	fmt.Println("")

	// teeReader -- T 形
	sr := strings.NewReader("1")
	buf = bytes.NewBuffer(src)
	tr := io.TeeReader(sr, buf) // 从 sr 读到的数据，在写入 buf 的同时，还可以从 tr 中读到
	trData := make([]byte, sr.Len())
	_, err = tr.Read(trData)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s, %s\n", trData, buf)

	// multiWriter
	var w1, w2, w3 bytes.Buffer
	mw := io.MultiWriter(&w1, &w2, &w3) // 往 mw 写入数据的同时，也会写入 w1, w2, w2
	sr = strings.NewReader("88")
	_, err = io.Copy(mw, sr)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s, %s, %s\n", w1.String(), w2.String(), w3.String())
}
