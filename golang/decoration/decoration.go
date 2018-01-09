// 修饰器
package main

import (
	"log"
	"time"
)

func main() {
	logTime(sleep10)
}

func sleep10() {
	time.Sleep(10 * time.Second)
}

func logTime(f func()) {
	bt := time.Now().Unix()
	f()
	et := time.Now().Unix()
	log.Printf("use time: %v", et-bt)
}
