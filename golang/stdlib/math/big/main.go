package main

import (
	"log"
	"math/big"
)

func main() {
	log.Println(big.MinExp, big.MaxExp, big.MaxPrec, big.MaxBase)

	for _, tc := range []struct {
		ix int64
		iy int64
		iz int
	}{
		{1, 1, 1},
		{0, 1, 1},
		{-1, 1, 1},
		{1, 3, 1},
		{0, 3, 0},
		{-1, 3, -1},
	} {
		x, y := big.NewInt(tc.ix), big.NewInt(tc.iy)
		iz := big.Jacobi(x, y)
		if iz != tc.iz {
			log.Fatalf("not equal, want %d, have %d", tc.iz, iz)
		}
		log.Println(iz)
	}
}
