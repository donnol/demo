package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"
)

func main() {
	x := "(1+2)*3"
	expr, err := parser.ParseExpr(x)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", expr)

	fset := &token.FileSet{}
	fn := "main.go"
	src := "package main"
	f, err := parser.ParseFile(fset, fn, src, parser.PackageClauseOnly)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v, %+v\n", f, fset)

	filter := func(os.FileInfo) bool {
		return true
	}
	path := "."
	pkgs, err := parser.ParseDir(fset, path, filter, parser.PackageClauseOnly)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v, %+v\n", pkgs, fset)

	src = "1+2"
	e, err := parser.ParseExprFrom(fset, fn, src, parser.PackageClauseOnly)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v, %+v\n", e, fset)
}
