// From https://making.pusher.com/golangs-real-time-gc-in-theory-and-practice/
// https://dave.cheney.net/2014/07/11/visualising-the-go-garbage-collector
// 	gcvis go run main.go
// https://blog.golang.org/go15gc
// https://talks.golang.org/2015/go-gc.pdf
package main

import (
	"fmt"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type (
	message []byte
	buffer  [windowSize]message
)

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(b *buffer, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	for {
		var b buffer
		for i := 0; i < msgCount; i++ {
			pushMsg(&b, i)
		}
		fmt.Println("Worst push time: ", worst)

	}
}
