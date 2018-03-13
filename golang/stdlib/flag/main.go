package main

import (
	"flag"
	"log"
)

func main() {
	var p string
	flag.StringVar(&p, "p", "default", "")
	flag.Parse()
	log.Printf("%s\n", p)

	// arg := flag.Arg(1)
	// log.Printf("%v\n", arg)
}
