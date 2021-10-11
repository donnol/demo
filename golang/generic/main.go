// From https://colobu.com/2021/03/22/try-go-generic/
// 
// more example: https://github.com/mattn/go-generics-example

package main

import (
    "fmt"

    "./tool"
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

type I0 interface {
    Name() string
}

type I1 interface {
    Age() int
}

type I interface {
    I0
    I1
}

type impl struct {

}

func (impl *impl) Name() string {
    return "impl"
}

func (impl *impl) Age() int {
    return 11
}

var _ I = (*impl)(nil)

func PrintI[T I](i T) {
    fmt.Println(i.Name(), i.Age())
}

type I3 interface {
    ~int | ~string
    I
}

func PrintI3[T I3](i T) {
    // 现在的type switch，还需要先把类型先转为interface{}
    // https://github.com/golang/go/issues/45380
    // 如果实现了这个提案，则不需要转为interface{}
    // 直接根据类型参数T的具体类型来决定：
    // switch type T {
    // case ~int & I:
    // case ~string & I:
    // }
    switch ii := interface{}(i).(type) {
    case int:
        fmt.Println(ii)
    case string:
        fmt.Println(ii)
    case I:
        fmt.Println(ii.Name(), ii.Age())
    }
}

type impl2 int

func (impl impl2) Name() string {
    return "impl2"
}

func (impl impl2) Age() int {
    return 11
}

// 满足条件：union element `~int | ~int8 | ~int16 | ~int32 | ~int64`中的一个，并且实现了`String() string`方法
type StringableSignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
	String() string
}

type SS int

func (ss SS) String() string {
    return "ss"
}

func PrintSSI[T StringableSignedInteger](i T) {
    fmt.Println(i)
}

// func div(a, b Addable) Addable {
//     return a/b
// }

// https://github.com/golang/go/issues/48424
// 根据这个提案，each函数的签名可以改为以下这种方式：
// 也就是省略了interface，直接使用了map[K]V, string等类型来做约束，非常方便
// func each[T interface{map[K]V}, K comparable, V interface{string}](m T) {
func each[T map[K]V, K comparable, V ~string|~int](m T) {
    for k, v := range m {
        fmt.Println(k, v)
    }
}

// string和~string
// string表示只适用于string类型
// ~string表示除了string类型，以string类型为底层类型的也一样适用，如：type MyString string里的MyString

// 如果把string|int从泛型类型约束里解放出来，支持下面这种写法，跟typescript不就一个样了？
// type Sai = string | int

type M[T any] struct {
    attr T
}

func (m M[T]) Attr() T {
    return m.attr
}

func (m *M[T]) Set(t T) {
    m.attr = t
}

// invalid AST: method must have no type parameters
// func (m *M[T]) Invalid[K any](k K) {
//     fmt.Println(k)
// }

type MM int

// invalid AST: method must have no type parameters
// func (m MM) Value[K int]() K {
//     return int(m)
// }

func Value[K ~int](m MM) K {
    // return int(m) // incompatible type: cannot use int(m) (value of type int) as K value
    return 0
}

func Value2[K ~int](k K) MM {
    return MM(k)
}

func Pointer[K ~int](k K) *K {
    return &k
}

func Pointer2[K ~int](k K) *int {
    var r = int(k)
    return &r
}

func Pointer3[K ~int](k K) *MM {
    var r = MM(k)
    return &r
}

func main() {
    // why not use <> instead [] in type parameter
    // because below code is valid before generic
    // so, in order to keep the compiler simple.
    a, b := 1 < 2, 2 > 3
    fmt.Println(a, b)

    fmt.Println(add(1,2))

    fmt.Println(add("foo","bar"))

    impl := &impl{}
    PrintI(impl)

    // var i I3 // interface contains type constraints
    // i = 1
    // fmt.Println(i)

    // PrintSSI(1) // int does not satisfy StringableSignedInteger (missing method String)
    PrintSSI(SS(1))

    PrintI3(impl2(1))

    m := make(map[int]string)
    m[1] = "jd"
    m[2] = "dd"
    each(m)

    m2 := make(map[int]int)
    m2[1] = 1
    m2[2] = 2
    each(m2)

    type MyString string
    m3 := make(map[int]MyString)
    m3[1] = "hah"
    m3[2] = "bab"
    each(m3)

    mi := M[int]{attr: 1}
    fmt.Println(mi.Attr())
    ms := M[string]{attr: "abc"}
    fmt.Println(ms.Attr())
    (&ms).Set("efg")
    fmt.Println(ms.Attr())

    mbm := tool.Map([]string{"jd", "jc"})
    fmt.Println(mbm)
}
