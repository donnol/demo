package main

import (
	"os"
)

func main() {
	file := "about.md"
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	for i := 0; i < 1000; i++ {
		_, err = f.WriteString("1")
		if err != nil {
			panic(err)
		}
	}
}
