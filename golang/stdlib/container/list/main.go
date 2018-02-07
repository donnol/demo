package main

import (
	"container/list"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	listTest()
}

func listTest() {
	l := list.New()
	l.Init()

	e := l.Front()
	log.Println(e)

	e = l.Back()
	log.Println(e)

	for _, v := range []int{
		1,
		2,
		3,
	} {
		e = l.PushBack(v)
		log.Println(e)
	}
	log.Printf("%+v", l)

	// 前序遍历
	for e := l.Front(); e != nil; e = e.Next() {
		log.Println(e)
	}

	// 后序遍历
	for e := l.Back(); e != nil; e = e.Prev() {
		log.Println(e)
	}

	// 取出
	e = l.Front()
	log.Println(e)

	e = l.Back()
	log.Println(e)

	log.Printf("%+v", l)

	// 逐个移除
	for e := l.Back(); e != nil; {
		e1 := e.Prev()
		v := l.Remove(e)
		log.Printf("%+v, %+v", l, v)
		e = e1
	}
}
