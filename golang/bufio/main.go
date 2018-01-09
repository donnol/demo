// reader.ReadLine 会使用 bufio.Reader.buf 字段，因此在调用多一次 ReadLine 后，前面的 a 变量会被后面的 b 变量覆盖
// 	可参看 bufio.Reader 里面 ReadLine 和 ReadSlice 方法的实现
// 	https://golang.org/src/bufio/bufio.go?s=9426:9493#L366
// 	https://golang.org/src/bufio/bufio.go?s=9426:9493#L313
// reader.ReadBytes 会 make 一个新的数组
package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	// readLine()

	readBytes()
}

func readLine() {
	reader := bufio.NewReader(os.Stdin)
	a, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", a)
	b, _, err := reader.ReadLine()
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", a) // 变成和 b 一样了
	log.Printf("%s\n", b)
}

func readBytes() {
	reader := bufio.NewReader(os.Stdin)
	a, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", a)
	b, err := reader.ReadBytes('\n')
	if err != nil {
		panic(err)
	}
	log.Printf("%s\n", a)
	log.Printf("%s\n", b)
}
