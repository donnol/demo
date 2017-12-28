package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	p "demo/morning-jog/2017-12-01"
	_ "demo/morning-jog/2017-12-28"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	testSort()
}

func testSort() {
	wanted, testCase := makeSortCase(100, 100)

	var worst time.Duration
	for _, tc := range testCase {
		// ut := wrap(p.Sort, tc, wanted) // 1000, 100000 14.582ms; 10, 1000000 78.041ms; 10, 10000000 913.975ms; 5, 100000000 10.565506s
		ut := wrap(p.SortConcurrent, tc, wanted) // 1000, 100000 28.035ms; 10, 1000000 101.013ms; 10, 10000000 987.183ms; 5, 100000000 10.572266s
		if ut > worst {
			worst = ut
		}
	}

	fmt.Printf("worst time: %v\n", worst)
}

func wrap(f func([]int) []int, had, wanted []int) time.Duration {
	bt := time.Now()
	r := f(had)
	ut := time.Since(bt)

	if !reflect.DeepEqual(r, wanted) {
		fmt.Printf("sort failed, get: %v, used time: %v\n", r, ut)
	} else {
		// fmt.Printf("used time: %v\n", ut)
	}
	return ut
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
