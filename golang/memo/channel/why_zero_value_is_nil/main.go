// 为什么 chan 的零值是 nil ?
// https://medium.com/justforfunc/why-are-there-nil-channels-in-go-9877cc0b2308
// 结合 空对象模式 理解，让零值也有意义
// 源码：https://github.com/campoy/justforfunc/blob/master/26-nil-chans/main.go
package main

import (
	"fmt"
	"time"
)

func main() {
	// var c chan int

	// fmt.Println(c) // <nil>

	// read -- block forever until 'fatal error: all goroutines are asleep -  deadlock'
	// <-c

	// write -- same as read
	// c <- 1

	// close -- panic: close of nil channel
	// close(c)

	// merge
	var a, b, c = make(chan int), make(chan int), make(chan int)

	// 从 c 读取数据
	go func() {
		for v := range c {
			fmt.Println(v)
		}
	}()

	// 将 a b 数据汇入 c
	go func() {
		// merge(c, a, b)
		for a != nil || b != nil {
			fmt.Println("receive and send", a, b)
			select {
			case v, ok := <-a:
				if ok {
					c <- v
				}
			case v, ok := <-b:
				if ok {
					c <- v
				}
			}
		}

		fmt.Println("close c chan")
		close(c)
	}()

	// 分别往 a b 写数据
	go func() {
		for a != nil {
			fmt.Println(a)
			a <- 0
		}
	}()
	go func() {
		for b != nil {
			b <- 1
		}
	}()

	// 两秒后停止
	after := time.After(time.Second * 2)
	select {
	case <-after:
		fmt.Println("==================================")
		a = nil
		b = nil
		fmt.Println(a, b)
		break
	}

	time.Sleep(time.Second * 2)
}

// Every element received in a or b will be sent to c, and once both a and b are closed c will be closed too.
func merge(c chan<- int, a, b <-chan int) { // 此处的 a b 与 main 里的 a b 不是同一个变量，只是它们的值里都含有相同的指针，指向同一块内存；当 main 里的 a b 被置为 nil 之后，这里的 a b 并没有改变
	for a != nil || b != nil {
		fmt.Println("receive and send", a, b)
		select {
		case v, ok := <-a:
			if ok {
				c <- v
			}
		case v, ok := <-b:
			if ok {
				c <- v
			}
		}
	}
	fmt.Println("merge finish")
}
