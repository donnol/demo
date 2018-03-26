package main

import (
	"go/printer"
	"go/token"
	"log"
	"os"
)

func main() {
	fset := &token.FileSet{}
	node := printer.CommentedNode{}
	err := printer.Fprint(os.Stdout, fset, node)
	if err != nil {
		log.Fatal(err) // unsupported node type printer.CommentedNode
	}
}
