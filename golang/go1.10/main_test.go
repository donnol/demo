package main

import (
	"fmt"
	"testing"
)

func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%s\n", "hello")
	}
}

func BenchmarkShuffle(b *testing.B) {
	var s []int
	for i := 0; i < b.N; i++ {
		s = append(s, i)
		shuffle(s)
	}
}
