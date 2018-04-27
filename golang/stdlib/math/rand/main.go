package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	log.Println(rand.ExpFloat64())  // 范围内呈指数分布的 float64 值
	log.Println(rand.NormFloat64()) // 范围内呈正态分布的 float64 值

	for i := 0; i < 10; i++ {
		s := rand.Perm(i)
		log.Println(s) // 一个 i 个值的切片，来自默认Source的整数[0，i）的伪随机置换

		p := make([]byte, i)
		n, err := rand.Read(p)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(n, string(p), p)

		// 洗牌
		rand.Shuffle(len(s), func(i int, j int) {
			s[i], s[j] = s[j], s[i]
		})
		log.Println(s)
	}

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := 0; i < 10; i++ {
		log.Println(r.Intn(i + 1))
	}
}
