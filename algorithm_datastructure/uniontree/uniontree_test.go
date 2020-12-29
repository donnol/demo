package uniontree

import "testing"

func TestUnionTree(t *testing.T) {
	/*
	   [[1,1,0],[1,1,0],[0,0,1]]
	   [[1]]
	   [[1,0,0],[0,1,0],[0,0,1]]
	   [[1,0,1],[0,1,0],[1,0,1]]
	   [[1,0,1,0],[0,1,0,0],[1,0,1,1],[0,0,1,1]]
	*/
	for i, cas := range []struct {
		m    [][]int
		want int
	}{
		{
			m:    [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
			want: 2,
		},
		{
			m:    [][]int{{1}},
			want: 1,
		},
		{
			m:    [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
			want: 3,
		},
		{
			m:    [][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}},
			want: 2,
		},
		{
			m:    [][]int{{1, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 1, 1}, {0, 0, 1, 1}},
			want: 2,
		},
	} {
		r := FindCircleNum(cas.m)
		if r != cas.want {
			t.Fatalf("No.%d, Bad result: %v != %v\n", i, r, cas.want)
		}
	}
}

func BenchmarkUnionTree(b *testing.B) {
	m := [][]int{{1, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 1, 1}, {0, 0, 1, 1}}
	for i := 0; i < b.N; i++ {
		FindCircleNum(m)
	}
}
