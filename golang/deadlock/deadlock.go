package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 0)
	// go func() {
	// 	ch <- 0
	// }()
	go func() {
		fmt.Println(<-ch)
		fmt.Println("hello")
	}()
	ch <- 0
	// fmt.Println(<-ch)
	time.Sleep(time.Second)
}
