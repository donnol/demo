// go1.10的变动
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	fmt.Println("helo go1.10")

	// 新建 strings.Builder
	// b := strings.Builder{}
	var b strings.Builder

	// 查看长度和内容
	fmt.Println("define:", b.Len(), b.String())

	for i, data := range []string{
		"hello",
		" ",
		"go1.10",
	} {
		// 扩容，长度还是不变
		b.Grow(len(data))
		fmt.Println("after grow:", b.Len(), b.String())

		if i%2 == 0 {
			// 写入 bytes
			n, err := b.Write([]byte(data))
			if err != nil {
				panic(err)
			}

			fmt.Println("after write:", b.Len(), b.String(), n)
		} else {
			// 写入字符串
			n, err := b.WriteString(data)
			if err != nil {
				panic(err)
			}

			fmt.Println("after WriteString:", b.Len(), b.String(), n)
		}

		// 重置
		// b.Reset()
		// fmt.Println("after reset:", b.Len(), b.String())
	}

	// 非常量位的常量操作， a 是变量
	// 未指定类型常量，1.0
	var a uint = 1
	var s = []int{1, 2, 3}
	fmt.Println(s[1.0<<a]) // 作为切片的下标时，1.0 是整型类型
	// fmt.Println(1.0 << a)  // 不作为下标时，1.0 是浮点数类型

	fmt.Println(M{"jd", 29}.String())

	shuffle(s)
	fmt.Println(s)
}

// M M
type M struct {
	Name string
	age  int
}

func (m M) String() string {
	return fmt.Sprintf("name is %s", m.Name)
}

func shuffle(s []int) {
	rand.Shuffle(len(s), func(i int, j int) {
		if s[i] > s[j] {
			s[i], s[j] = s[j], s[i]
		}
	})
}
