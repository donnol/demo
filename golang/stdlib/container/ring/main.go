package main

import (
	"container/ring"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	ringTest()
}

func ringTest() {
	l := 2

	// 新建
	r := ring.New(l) // 指定长度
	r.Value = l
	log.Printf("%+v, %+v, %v", r, r.Value, r.Len() == l)

	// 下一个
	r = r.Next()
	r.Value = l + 1
	log.Printf("%+v, %+v, %v", r, r.Value, r.Len() == l)

	// 上一个
	r = r.Prev()
	log.Printf("%+v, %+v, %v", r, r.Value, r.Len() == l)

	// 向前移动
	r = r.Move(l)
	log.Printf("%+v, %+v, %v", r, r.Value, r.Len() == l)

	// 打印每个元素
	r.Do(func(v interface{}) {
		log.Printf("%+v", v)
	})

	// 再建一个
	nr := ring.New(l + 1)
	nr.Do(func(v interface{}) {
		log.Printf("%+v", v)
	})
	for i := 0; i < nr.Len(); i++ {
		nr.Value = i + 1
		nr = nr.Next()
	}
	nr.Do(func(v interface{}) {
		log.Printf("%+v", v)
	})

	// 连接两个环
	ar := r.Link(nr)
	ar.Do(func(v interface{}) {
		log.Printf("%+v", v)
	})

	// 从环里去掉 l 个，从 ar.Next() 开始
	ar.Unlink(l)
	ar.Do(func(v interface{}) {
		log.Printf("%+v", v)
	})
}
