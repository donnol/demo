package main

import (
	"go/scanner"
	"os"
)

type myError struct {
}

func (m myError) Error() string {
	return "my error"
}

func main() {
	err := myError{}
	scanner.PrintError(os.Stdout, err) // my error
}
