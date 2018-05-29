package binarysearchtree

import (
	"math/rand"
	"testing"
	"time"
)

var s = []string{"a", "b", "c", "d", "e", "f", "g", "h"}

func init() {
	rand.Seed(time.Now().Unix())

	shuffle(s)
}

func TestBinarySearchTree(t *testing.T) {
	// insert
	bst := &BinarySearchTree{}
	for _, key := range s {
		value := key
		bst.Insert(key, value)
	}

	// string
	bst.String()
}

func shuffle(s []string) {
	rand.Shuffle(len(s), func(i int, j int) {
		if s[i] > s[j] {
			s[i], s[j] = s[j], s[i]
		}
	})
}

func BenchmarkBinarySearchTreeInsert(b *testing.B) {
	bst := &BinarySearchTree{}
	for i := 0; i < b.N; i++ {
		key := s[0]
		value := key
		bst.Insert(key, value)
	}
}
