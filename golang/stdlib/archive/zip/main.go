package main

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	filename := "test.zip"
	filename2 := "create.zip"
	read(filename)
	write(filename2)
	header(filename)
}

// 打开 zip 文件，并读取文件里的内容
func read(filename string) {
	rc, err := zip.OpenReader(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer rc.Close()

	for _, file := range rc.File {
		f, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		data, err := ioutil.ReadAll(f)
		if err != nil {
			log.Fatal(err)
		}

		log.Println(string(data))
	}
}

// 将内容写入 zip 中
func write(filename string) {
	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)
	for _, file := range []struct {
		Name, Body string
	}{
		{"a.txt", strings.Repeat("a", 1000) + "\n"},
		{"b.txt", strings.Repeat("b", 1000) + "\n"},
	} {
		fh := &zip.FileHeader{
			Name: file.Name,
			// Comment: file.Body,
			CompressedSize64:   5,
			UncompressedSize64: 5,
		}
		iow, err := w.CreateHeader(fh)
		if err != nil {
			log.Fatal(err)
		}
		_, err = iow.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Close()

	err := ioutil.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// zip 文件的基本信息
func header(filename string) {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fih, err := zip.FileInfoHeader(fi)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", fih)
}
