package main

import (
	"go/ast"
	"log"
)

func main() {
	file := &ast.File{}
	export := ast.FileExports(file)
	log.Println(export)

	// ast.FilterDecl(decl, ast.Filter)

	plus := ast.NewIdent("+")

	log.Printf("%+v, %v, %v, %v\n", plus, plus.Pos(), plus.End(), plus.IsExported())
}
