package main

import (
	"fmt"
	"log"
)

func main() {
	data, err := Asset("about.html")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
