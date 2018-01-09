package main

import (
	"os"
	"runtime/trace"
)

func main() {
	trace.Start(os.Stderr)
	ch := make(chan int)
	go func() {
		ch <- 42
	}()
	<-ch
	trace.Stop()
}
