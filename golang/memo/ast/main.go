package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	// testExpr()

	// testFileExports()

	testInspect()

	testWalk()
}

// 测试表达式
func testExpr() {
	testCase := []string{
		`a == 1`,
		`b == 'b'`,
		`a == func() int {return 1}()`,
		`1 + 1`,
		`1 == "1"`,

		`s[0]`,
		`m["k"]`,

		`f(1)`,
		`f(b...)`,

		`m{"k":1}`,

		`(1==1)`,

		`p.v`,

		`s[:]`,

		`*p`,

		`i.(type)`,
		`i.(int)`,

		`!b`,
	}

	for _, tc := range testCase {
		expr, err := parser.ParseExpr(tc)
		if err != nil {
			log.Fatal(err)
		}

		switch e := expr.(type) {
		case *ast.BinaryExpr:
			log.Printf("BinaryExpr is: %+v, X: %+v, Y: %+v\n", e, e.X, e.Y)
		case *ast.IndexExpr:
			log.Printf("IndexExpr is: %+v, %+v, %+v, %+v, %+v\n", e, e.X, e.Index, e.Lbrack, e.Rbrack)
		case *ast.CallExpr:
			log.Printf("CallExpr is: %+v, %+v, %+v, %+v, %+v, %+v\n", e, e.Args, e.Ellipsis, e.Fun, e.Lparen, e.Rparen)
		case *ast.CompositeLit:
			for _, se := range e.Elts {
				switch es := se.(type) {
				case *ast.KeyValueExpr:
					log.Printf("KeyValueExpr is: %+v, %#v, %#v", es, es.Key, es.Value)
				default:
					log.Printf("CompositeLit is: %+v, %#v\n", e, e.Elts)
				}
			}
		case *ast.ParenExpr:
			log.Printf("ParenExpr is: %+v, %+v\n", e, e.X)
		case *ast.SelectorExpr:
			log.Printf("SelectorExpr is: %+v, %+v, %+v\n", e, e.Sel, e.X)
		case *ast.SliceExpr:
			log.Printf("SliceExpr is: %+v\n", e)
		case *ast.StarExpr:
			log.Printf("StarExpr is: %+v\n", e)
		case *ast.TypeAssertExpr:
			log.Printf("TypeAssertExpr is: %+v\n", e)
		case *ast.UnaryExpr:
			log.Printf("UnaryExpr is: %+v\n", e)
		default:
			log.Printf("default: %#v\n", e)
		}
	}
}

var testCase = []string{
	`package p
	const c = 1.0
	var X = f(3.14)*2 + c
	`,
	// `package p
	// const a = 2
	// var x = a+1
	// `,
	// `package p
	// func a() {
	// 	a++
	// }
	// `,
	// `package p
	// func B() {
	// 	b++
	// }
	// `,
}

func testFileExports() {
	for _, tc := range testCase {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "src.go", tc, 0)
		if err != nil {
			panic(err)
		}
		if ast.FileExports(file) {
			log.Printf("true")
		} else {
			log.Printf("false")
		}
	}
}

func testInspect() {
	for _, tc := range testCase {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "src.go", tc, 0)
		if err != nil {
			panic(err)
		}
		ast.Inspect(file, func(n ast.Node) bool {
			if n != nil {
				log.Printf("=== inspect: %#v\n", n)
			}
			return true
		})
		break
	}
}

type visit struct {
}

func (v visit) Visit(node ast.Node) (w ast.Visitor) {
	if node != nil {
		log.Printf("node is: %#v\n", node)
	}

	return v
}

func testWalk() {
	for _, tc := range testCase {
		fset := token.NewFileSet()
		file, err := parser.ParseFile(fset, "src.go", tc, 0)
		if err != nil {
			panic(err)
		}
		ast.Walk(visit{}, file)
	}
}
