// 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数
package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	r := getZeroWithoutReturn()
	if r <= 0 {
		log.Fatal("failed.")
	} else {
		log.Println("successed, get", r)
	}
}

func getZeroWithoutReturn() (r int) {
	v := rand.Intn(100) + 1
	defer func() {
		if e := recover(); e != nil {
			r = v
		}
	}()
	panic(v)
}
