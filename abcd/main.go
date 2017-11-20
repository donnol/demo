package main

import (
	"fmt"
	"os"
	"time"
)

var (
	count int
	ch    chan Model
)

// 将要写入文件的数字
type Model struct {
	File *os.File
	Num  int
}

func init() {
	count = 4
	ch = make(chan Model)
}

func main() {
	ch1 := gen(1)
	ch2 := gen(3)
	ch3 := gen(5)
	ch4 := gen(7)

	var fs []*os.File
	for _, fileName := range []string{"A", "D", "C", "B"} {
		f, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			panic(err)
		}
		fs = append(fs, f)
	}

	go write()

	timeCh := time.After(time.Millisecond * 10)
	times := 0
	for {
		select {
		case <-timeCh:
			return
		default:
			n := <-ch1
			index := times % count
			ch <- Model{fs[index], n}

			n = <-ch2
			index = times%count - 1
			if index < 0 {
				index += count
			}
			if index > count-1 {
				index -= count
			}
			ch <- Model{fs[index], n}

			n = <-ch3
			index = times%count - 2
			if index < 0 {
				index += count
			}
			if index > count-1 {
				index -= count
			}
			ch <- Model{fs[index], n}

			n = <-ch4
			index = times%count - 3
			if index < 0 {
				index += count
			}
			if index > count-1 {
				index -= count
			}
			ch <- Model{fs[index], n}
			times++
		}
	}
}

func write() {
	for m := range ch {
		_, err := fmt.Fprintf(m.File, "%d", m.Num)
		if err != nil {
			panic(err)
		}
	}
}

func gen(num int) chan int {
	ch := make(chan int)
	go func() {
		for {
			ch <- num
		}
	}()
	return ch
}
