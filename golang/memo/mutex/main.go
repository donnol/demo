package main

import (
	"sync"
)

func main() {
	f := func() {
		panic("oh my god")
		// os.Exit(1)
	}
	withMutex(f)
}

var mutex sync.Mutex

func withMutex(f func()) {
	mutex.Lock()
	defer mutex.Unlock()
	f()
	// mutex.Unlock()
	// log.Println("unlock")
}
