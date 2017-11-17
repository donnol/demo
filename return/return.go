package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(ret())
}

func ret() (a int, err error) {
	a = 1
	err = errors.New("abc")
	return err
	// return
}
