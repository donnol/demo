package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 随机生成fib数字
	chGen := make(chan int, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			chGen <- genFib()
		}
		close(chGen)
	}()

	// 计算和
	chResult := make(chan int, 1)
	go func() {
		sum := 0
		for v := range chGen {
			sum += v
		}
		chResult <- sum
	}()

	// 打印结果
	// select {
	// case <-time.After(time.Second * 30):
	// 	close(chGen)
	// }

	sum := <-chResult
	fmt.Println(sum)
}

var m map[int]int

func init() {
	m = make(map[int]int)
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	if v, ok := m[n]; ok {
		return v
	}
	r := fib(n-2) + fib(n-1)
	m[n] = r
	return r
}

func genFib() int {
	n := rand.Intn(100)
	return fib(n)
}
