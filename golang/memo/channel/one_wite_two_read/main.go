// 一个管道写，两个管道读，谁先读到呢
// 看标准库 database/sql sql.go
//  多个地方使用了 <-tx.ctx.Done()
// 生此疑问
package main

import (
	"log"
	"time"
)

func main() {
	// 打印文件行号
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	done := make(chan int)
	// done := make(chan int, 1)

	go func() {
		<-done
		log.Printf("%d\n", 2)
	}()
	go func() {
		<-done
		log.Printf("%d\n", 1)
	}()

	done <- 1
	done <- 2

	time.Sleep(time.Second)
}
