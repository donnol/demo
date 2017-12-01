package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	p "demo/morning-jog/2017-12-01"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	testSort()
}

func testSort() {
	wanted, testCase := makeSortCase(1000, 1000000)

	var worst time.Duration
	for _, tc := range testCase {
		bt := time.Now()
		r := p.Sort(tc)
		ut := time.Since(bt)
		if ut > worst {
			worst = ut
		}
		if !reflect.DeepEqual(r, wanted) {
			fmt.Printf("sort failed, get: %v\n", r)
		} else {
			fmt.Printf("used time: %v\n", ut)
		}
	}
	fmt.Printf("worst time: %v\n", worst)
}

// 生成cnt个从0到num-1的随机数组
func makeSortCase(cnt, num int) (sorted []int, randoms [][]int) {
	// 顺序数组
	for i := 0; i < num; i++ {
		sorted = append(sorted, i)
	}

	for i := 0; i < cnt; i++ {
		random := shuffle(sorted)
		randoms = append(randoms, random)
	}

	return
}

// 洗牌
func shuffle(s []int) []int {
	bt := time.Now()
	sc := shuffle2(s)
	ut := time.Since(bt)
	fmt.Printf("shuffle used time: %v\n", ut)
	// fmt.Println(s, sc)

	return sc
}

// Knuth-Durstenfeld Shuffle -- O(n) in-place(改变原数组)
func shuffle1(sc []int) {
	l := len(sc)

	// 逆向
	// for i := (l - 1); i >= 1; i-- {
	// 	randNum := rand.Intn(i + 1)
	// 	sc[i], sc[randNum] = sc[randNum], sc[i]
	// }

	// 顺向
	for i := 0; i <= (l - 1); i++ {
		randNum := rand.Intn(l-i) + i
		sc[i], sc[randNum] = sc[randNum], sc[i]
	}
}

// Inside-Out Algorithm -- 不改变原数组
func shuffle2(source []int) []int {
	l := len(source)

	// 新建数组
	var sc = make([]int, l)
	copy(sc, source) // sc 的长度要与 s 相同，否则无法复制完整内容

	for i := 1; i < l; i++ {
		randNum := rand.Intn(i + 1)
		if randNum != i {
			sc[i] = sc[randNum]
		}
		sc[randNum] = source[i]
	}

	return sc
}
