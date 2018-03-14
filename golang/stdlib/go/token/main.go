package main

import (
	"go/token"
	"log"
)

func main() {
	fs := token.NewFileSet()
	log.Printf("%+v, %d\n", fs, fs.Base())

	f := fs.AddFile("hello.go", fs.Base(), 0)
	log.Printf("%+v, %+v\n", fs, f)

	f = fs.File(f.Pos(0))
	log.Printf("%+v, %+v\n", fs, f)
	log.Printf("%d, %d, %d, %s, %d", f.Size(), f.Base(), f.Line(f.Pos(0)), f.Name(), f.LineCount())

	fs.AddFile("world.go", fs.Base(), 0)

	fs.Iterate(func(f *token.File) bool {
		log.Printf("%+v\n", f)
		return true // 返回 true 继续执行下一个，返回 false 则退出
	})

	log.Printf("%v\n", f.Pos(0).IsValid())

	pos := f.Position(f.Pos(0))
	log.Printf("%v, %v\n", pos.String(), pos.IsValid())

	for _, ident := range []string{
		"var a int",
		"break",
		`if true { 
			break
		}`,
		"a = 1",
		"123.4",
		"1+1",
		"+",
	} {
		t := token.Lookup(ident)
		log.Printf("%+v, %v, %v, %v, %v\n", t, t.IsKeyword(), t.IsLiteral(), t.IsOperator(), t.Precedence())
	}
}
