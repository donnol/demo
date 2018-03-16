package main

import (
	"go/constant"
	"go/token"
	"log"
)

func main() {
	x := constant.MakeInt64(60)
	bytes := constant.Bytes(x)
	log.Printf("%v, %s\n", bytes, bytes)

	bv := constant.MakeBool(true)
	b := constant.BoolVal(bv)
	log.Printf("%v, %v\n", bv, b)

	log.Printf("%v, %v, %v\n", bv.Kind(), bv.String(), bv.ExactString())

	op := token.ADD
	y := constant.MakeInt64(10)
	v := constant.BinaryOp(x, op, y)
	log.Printf("%v\n", v)
}
