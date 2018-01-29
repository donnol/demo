package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	fileInfo()
	create()
	read()
}

func fileInfo() {
	fi, err := os.Stat("test.tar")
	if err != nil {
		log.Fatal(err)
	}
	h, err := tar.FileInfoHeader(fi, "test.tar")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", h)
}

// 创建 tar 档案
func create() {
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	for _, file := range []struct {
		Name, Body string
	}{
		{"a.txt", "a\n"},
		{"b.txt", "b\n"},
	} {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0644,
			Size: int64(len(file.Body)),
		}
		err := tw.WriteHeader(hdr)
		if err != nil {
			log.Fatal(err)
		}
		_, err = tw.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	tw.Close()

	err := ioutil.WriteFile("create.tar", buf.Bytes(), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// 读取 tar 档案
func read() {
	// 可以使用 create 函数创建的 tar 文档
	// 也可以使用系统命令 tar -cf test.tar a.txt b.txt 生成的 tar 文档
	//  注意，使用 tar 命令时不要加 z 选项，不然会压缩成 zip 文件，这样就读取不了了
	file := "create.tar"
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	buf := bytes.NewBuffer(data)
	tr := tar.NewReader(buf)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(hdr.Name)

		var trBuf = make([]byte, hdr.Size)
		_, err = tr.Read(trBuf)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(trBuf))
	}
}
