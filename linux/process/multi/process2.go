package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	filename := "data.md"
	fi, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file size is %d\n", fi.Size())

	f, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b := []byte("b")
	dataLen := 0
	timer := time.NewTimer(5 * time.Second)
out:
	for {
		select {
		case <-timer.C:
			break out
		default:
			n, err := f.Write(b)
			if err != nil {
				log.Fatal(err)
			}
			dataLen += n
		}
	}
	fmt.Printf("write total length is %d\n", dataLen)
}
