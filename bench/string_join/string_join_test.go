package string_join

import "testing"

func BenchmarkStringPlus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringPlus()
	}
}

func BenchmarkBytesRead(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BytesRead()
	}
}

func BenchmarkAppend(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Append()
	}
}

func BenchmarkFmt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fmt()
	}
}
