package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	scanBytes(true)
	filename := "test.md"
	writer(filename)
	reader(filename)
	readWriter(filename)
}

func scanBytes(atEOF bool) {
	have := 0
	data := []byte("hello")
	// advance 表示拿了多少个，token 表示拿出来的字符
	for {
		advance, token, err := bufio.ScanBytes(data[have:], atEOF)
		if err != nil {
			log.Fatal(err)
		}
		if advance == 0 { // 没有字符拿到，表示已拿完
			break
		}
		have += advance
		log.Println(advance, string(token))
	}

	// 只有这个会因为 atEOF 的值不同而输出不同
	have = 0
	for {
		advance, token, err := bufio.ScanLines(data[have:], atEOF)
		if err != nil {
			log.Fatal(err)
		}
		if advance == 0 {
			break
		}
		have += advance
		log.Println(advance, string(token))
	}

	have = 0
	for {
		data = []byte("温暖")
		advance, token, err := bufio.ScanRunes(data[have:], atEOF)
		if err != nil {
			log.Fatal(err)
		}
		if advance == 0 {
			break
		}
		have += advance
		log.Println(advance, string(token))
	}

	have = 0
	for {
		data = []byte("温暖 和煦")
		advance, token, err := bufio.ScanWords(data[have:], atEOF)
		if err != nil {
			log.Fatal(err)
		}
		if advance == 0 {
			break
		}
		have += advance
		log.Println(advance, string(token))
	}
}

func writer(filename string) {
	// f, err := os.Open(filename) // 会出现 bad file descriptor 错误
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	data := []byte("hello\n")
	n, err := bw.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n, bw.Buffered(), bw.Available())

	err = bw.Flush()
	if err != nil {
		log.Fatal(err)
	}
}

func reader(filename string) {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	buf := make([]byte, fi.Size())
	_, err = br.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buf))
}

func readWriter(filename string) {
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	br := bufio.NewReader(f)
	bw := bufio.NewWriter(f)
	brw := bufio.NewReadWriter(br, bw)

	data := []byte("hello\n")
	size := int(fi.Size()) + len(data)
	n, err := brw.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(n)
	err = brw.Flush()
	if err != nil {
		log.Fatal(err)
	}

	// 上面 写入成功了，但是这里读的时候报错 EOF 了
	buf := make([]byte, size)
	_, err = brw.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(buf))
}
