package main

import (
	"errors"
	"log"
)

func main() {
	err := errors.New("Hello, I am jd")
	if err != nil {
		log.Fatal(err)
	}
}
