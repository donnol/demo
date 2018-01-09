package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("hello")

	m := sync.Map{}
	fmt.Printf("%+v", m)

	now := time.Now()
	time.Sleep(time.Second)
	d := time.Since(now)
	fmt.Println(d)
}
