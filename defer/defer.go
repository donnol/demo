package main

import (
	"errors"
	"fmt"
)

func testDefer() (err error) {
	err = errors.New("123")
	// return makeErr()
	return
}

func makeErr() error {
	return errors.New("abc")
}

func main() {
	fmt.Println(testDefer())
}
