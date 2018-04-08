// From https://lingchao.xin/post/memory-leaking.html
package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	var c chan bool
	// jLeak(c)
	jGood(c)
}

// no.1
var s0 string // package level variable

func fLeak(s1 string) {
	// 假设 s1 是一个长度大于 50 的字符串.
	s0 = s1[:50]
	// 现在, s0 和 s1 共享相同的底层内存块.
	// s1 现在不存活了, 但是 s0 依然存活.
	// 尽管仅有 50 个字节在内存块中,
	// s0 仍旧存活的事实阻止了这 1M 字节的内存块被回收.
}

// =>
func fGood(s1 string) {
	var b strings.Builder
	b.Grow(50)
	b.WriteString(s1[:50])
	s0 = b.String()
	// b.Reset() // 如果 b 在其他地方会用到, 那么它必须在这里重置掉.
}

// no.2
var a0 []int

func gLeak(s1 []int) {
	// 假设 s1 的长度远远大于 30.
	a0 = s1[len(s1)-30:]
}

// =>
func gGood(s1 []int) {
	a0 = append([]int(nil), s1[len(s1)-30:]...)
}

// no.3
func hLeak() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	return s[1:3]
}

// =>
func hGood() []*int {
	s := []*int{new(int), new(int), new(int), new(int)}
	s1 := s[1:3]
	s[0] = nil
	s[len(s)-1] = nil
	return s1
}

// no.4
// func kLeak(c <-chan bool) { // don't use leading k in Go names; func kLeak should be leak
func jLeak(c <-chan bool) {
	s := make([]int64, 1e6)
	if <-c { // 如果 c 为 nil, 这里将永远阻塞
		_ = s
		// 使用 s, ...
	}
}

// =>
// func kGood(c <-chan bool) { // don't use leading k in Go names; func kGood should be good
func jGood(c <-chan bool) {
	if c == nil {
		return
	}
	// ...
}

// no.5
func mLeak() {
	type T struct {
		v [1 << 20]int
		t *T
	}

	var finalizer = func(t *T) {
		fmt.Println("finalizer called")
	}

	var x, y T

	// SetFinalizer 会使 x 逃逸到堆上.
	runtime.SetFinalizer(&x, finalizer)

	// 以下语句将导致 x 和 y 变得无法收集.
	x.t, y.t = &y, &x // y 也逃逸到了 堆上.
}
