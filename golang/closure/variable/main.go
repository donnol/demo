// 哪些类型的变量值会随时间改变呢
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("begin", intClosure, stringClosure, sliceClosure, mapClosure, syncMapClosure)
	time.Sleep(10 * time.Second)
	fmt.Println("end", intClosure, stringClosure, sliceClosure, mapClosure, syncMapClosure)
}

var intClosure = func() int {
	var a int
	refresh := func() {
		a++
		fmt.Println(time.Now(), a)
	}
	refresh()
	go func() {
		for range time.NewTicker(time.Second).C {
			refresh()
		}
	}()
	return a
}()

var stringClosure = func() string {
	var str string
	refresh := func() {
		str += "go"
		fmt.Println(time.Now(), str)
	}
	refresh()
	go func() {
		for range time.NewTicker(time.Second).C {
			refresh()
		}
	}()
	return str
}()

var sliceClosure = func() []int {
	var s []int
	// var s = make([]int, 20) // 长度够也不会变
	var i int
	refresh := func() {
		i++
		s = append(s, i)
		fmt.Println(time.Now(), s)
	}
	refresh()
	go func() {
		for range time.NewTicker(time.Second).C {
			refresh()
		}
	}()
	return s
}()

var mapClosure = func() map[int]int {
	var m = make(map[int]int)
	var i int
	refresh := func() {
		i++
		m[i] = i
		fmt.Println(time.Now(), m)
	}
	refresh()
	go func() {
		for range time.NewTicker(time.Second).C {
			refresh()
		}
	}()
	return m
}()

var syncMapClosure = func() *sync.Map {
	var sm = new(sync.Map)
	var i int
	refresh := func() {
		i++
		sm.Store(i, i)
		fmt.Println(time.Now(), sm)
	}
	refresh()
	go func() {
		for range time.NewTicker(time.Second).C {
			refresh()
		}
	}()
	return sm
}()
