package main

import "fmt"

const primeRK = 10

func main() {
	type data struct {
		S      string
		Substr string
	}
	testCase := []struct {
		Data data
		Want bool
	}{
		{data{"", ""}, true},
		{data{"abc", ""}, false},
		{data{"abc", "a"}, true},
		{data{"abc", "abc"}, true},
		{data{"efg", "a"}, false},
		{data{"efg", "abcd"}, false},
		{data{"efg", "abcdefg"}, false},
		{data{"abcdefg", "abcdefg"}, true},
	}

	for _, tc := range testCase {
		if Contains(tc.Data.S, tc.Data.Substr) != tc.Want {
			panic(fmt.Sprintf("s: %s, substr: %s, want: %v, have: %v", tc.Data.S, tc.Data.Substr, tc.Want, Contains(tc.Data.S, tc.Data.Substr)))
		}
	}
}

func Contains(s, substr string) bool {
	n := len(substr)
	switch {
	case n == 0:
		if substr == s {
			return true
		}
		return false
	case n == len(s):
		if substr == s {
			return true
		}
		return false
	case n > len(s):
		return false
	}

	// Rabin-Karp
	hashss, pow := hashStrRev(substr)
	fmt.Printf("%d, %d\n", hashss, pow)
	last := len(s) - n
	var h uint32
	for i := len(s) - 1; i >= last; i-- {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[last:] == substr {
		return true
	}
	for i := last - 1; i >= 0; i-- {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i+n])
		if h == hashss && s[i:i+n] == substr {
			return true
		}
	}
	return false
}

func hashStrRev(s string) (uint32, uint32) {
	hash := uint32(0)
	for i := len(s) - 1; i >= 0; i-- {
		hash = hash*primeRK + uint32(s[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(s); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}
