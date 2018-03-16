package main

import (
	"go/ast"
	"go/doc"
	"log"
)

func main() {
	pkg := &ast.Package{}
	docPkg := doc.New(pkg, ".", doc.AllDecls)
	log.Printf("%+v\n", docPkg)

	for _, tc := range []string{
		"var a = 1",
		"var a int",
		"// a a",
		"// a a, a",
		"// a a; a",
		"// a a\na",
		"// a a. a",
	} {
		log.Printf("%v\n", doc.IsPredeclared(tc))
		log.Printf("%v\n", doc.Synopsis(tc)) // 截取'.'前面的作为概要
	}

}
