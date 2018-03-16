package main

import (
	"go/importer"
	"io"
	"log"
)

func main() {
	imp := importer.Default()
	log.Printf("%+v\n", imp)

	// compiler := "gccgo"
	compiler := "gc"
	lookup := func(path string) (rc io.ReadCloser, err error) {
		return
	}
	imp = importer.For(compiler, lookup)
	log.Printf("%+v\n", imp)
}
