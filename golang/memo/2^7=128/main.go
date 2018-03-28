package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	z, ok := assert(2, 7)
	log.Printf("%v, %v\n", z, ok)

	r := findNum()
	for _, s := range r {
		log.Printf("%s\n", s)
	}
}

type rule struct {
	x, y, z int
}

func (r rule) String() string {
	return fmt.Sprintf("x: %d, y: %d, z: %d\n", r.x, r.y, r.z)
}

func assert(x, y int) (z int, ok bool) {
	z = 100 + x*10 + (y + 1)
	p := math.Pow(float64(x), float64(y))
	if int(p) != z {
		return
	}
	ok = true
	return
}

// 寻找类似 2^7=128 的数字
func findNum() (rules []rule) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if z, ok := assert(i, j); ok {
				rules = append(rules, rule{
					i,
					j,
					z,
				})
			}
		}
	}
	return
}
