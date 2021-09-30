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
// 根据这个提案，each函数的签名可以改为：
// func each[T map[K]V, K comparable, V string](m T) {
// 也就是省略了interface，直接使用了map[K]V, string等类型来做约束，非常方便
func each[T interface{map[K]V}, K comparable, V interface{string}](m T) {
    for k, v := range m {
        fmt.Println(k, v)
    }
}

func main() {
    fmt.Println(add(1,2))

    fmt.Println(add("foo","bar"))

    m := make(map[int]string)
    m[1]= "jd"
    each(m)
}
