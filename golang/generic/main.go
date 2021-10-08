// From https://colobu.com/2021/03/22/try-go-generic/
// 
// more example: https://github.com/mattn/go-generics-example

package main

import (
    "fmt"
)

type Addable interface {
    // 即将被type set取代
	// type int, int8, int16, int32, int64,
	// 	uint, uint8, uint16, uint32, uint64, uintptr,
	// 	float32, float64, complex64, complex128,
	// 	string

    // type set
    int | int8 | int16 | int32 | int64 | string
}

func add[T Addable](a, b T) T {
    return a + b
}

// https://github.com/golang/go/issues/48424
// 根据这个提案，each函数的签名可以改为以下这种方式：
// 也就是省略了interface，直接使用了map[K]V, string等类型来做约束，非常方便
// func each[T interface{map[K]V}, K comparable, V interface{string}](m T) {
func each[T map[K]V, K comparable, V string|int](m T) {
    for k, v := range m {
        fmt.Println(k, v)
    }
}

// 如果把string|int从泛型类型约束里解放出来，支持下面这种写法，跟typescript不就一个样了？
// type Sai = string | int

func main() {
    // why not use <> instead [] in type parameter
    // because below code is valid before generic
    // so, in order to keep the compiler simple.
    a, b := 1 < 2, 2 > 3
    fmt.Println(a, b)

    fmt.Println(add(1,2))

    fmt.Println(add("foo","bar"))

    m := make(map[int]string)
    m[1] = "jd"
    m[2] = "dd"
    each(m)

    m2 := make(map[int]int)
    m2[1] = 1
    m2[2] = 2
    each(m2)
}
