package main

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	// src := []byte("Hello, I am jd") // expected 1 expression (and 1 more errors)
	src := []byte("var a = 1; var b = 2")
	dst, err := format.Source(src)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v, %s\n", dst, dst)

	var buf = new(bytes.Buffer)
	fset := &token.FileSet{}
	// node := "1+2" // unsupported node type string
	node, err := parser.ParseExpr("(1+2)*3/4")
	if err != nil {
		log.Fatal(err)
	}
	err = format.Node(buf, fset, node)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", buf.String())
}
