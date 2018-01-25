package main

import (
	"fmt"
	"time"
)

func main() {
	var sendChan = make(chan int)
	var sendChan1 = make(chan int)
	var closeChan = make(chan struct{})

	// select 是这样工作的：
	//  	它会逐个计算case，找出其中成功的，然后随机挑选一个运行
	go func(sendChan, sendChan1 chan int, closeChan chan struct{}) {
		var i, j int
		for {
			select {
			case v := <-sendChan:
				i++
				fmt.Println("receive from chan", v)
			case v := <-sendChan1:
				j++
				fmt.Println("receive from chan1", v)
			case <-closeChan:
				fmt.Println("receive from closeChan", i, j)
				return
			}
		}
	}(sendChan, sendChan1, closeChan)

	go func() {
		for {
			sendChan <- 1
		}
	}()
	go func() {
		for {
			sendChan1 <- 2
		}
	}()

	time.Sleep(time.Second * 1)
	closeChan <- struct{}{}
}
