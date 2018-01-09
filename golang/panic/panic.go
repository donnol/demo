package main

import (
	"fmt"
	"sync"
)

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		return
	// 	}
	// }()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer func() {
			if err := recover(); err != nil {
				return
			}
		}()
		err := run()
		if err != nil {
			panic(err)
		}
		panic("panic")
	}()
	wg.Wait()
}

func run() error {
	fmt.Println("i am running.")
	return nil
}
