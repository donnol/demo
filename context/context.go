package main

import (
	"context"
	"fmt"
)

func main() {
	b1 := context.Background()
	b2 := context.Background()
	fmt.Println(&b1, &b2 /*b1.(*context.emptyCtx).String()*/)
}
