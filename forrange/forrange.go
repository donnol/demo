package main

import "fmt"
import "time"

func main() {
	forRangeInt(10)
	time.Sleep(time.Second)

	ch := make(chan int, 10)
	for i := 0; i < 10; i++ {
		ch <- i
	}

	forRangeChan(ch)
	close(ch)
	time.Sleep(time.Second)
}

func forRangeInt(n int) {
	for i := 0; i < n; i++ {
		go func(i int) {
			time.Sleep(time.Millisecond)
			fmt.Println(i)
		}(i)
	}
}

func forRangeChan(ch chan int) {
	for v := range ch {
		go func(v int) {
			time.Sleep(time.Millisecond)
			fmt.Println(v)
		}(v)
	}
}
