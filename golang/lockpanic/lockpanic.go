package main

import (
	"sync"
)

func main() {
	lockpanic(false)

	lockpanic(true)
}

var m sync.Mutex

func lockpanic(p bool) {
	m.Lock()
	if p {
		panic("hi")
	}
	m.Unlock()
}
