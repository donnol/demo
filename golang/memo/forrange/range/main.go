package main

import (
	"fmt"
)

func main() {
	v := []int{1, 2, 3}
	fmt.Printf("%p, %d, %d\n", v, len(v), cap(v))

	// 编译时，转换为以下代码
	// https://github.com/golang/gofrontend/blob/e387439bfd24d5e142874b8e68e7039f74c744d7/go/statements.cc#L5593
	// The loop we generate:
	//   for_temp := range
	//   len_temp := len(for_temp)
	//   for index_temp = 0; index_temp < len_temp; index_temp++ {
	//           value_temp = for_temp[index_temp]
	//           index = index_temp
	//           value = value_temp
	//           original body
	//   }
	//
	// Using for_temp means that we don't need to check bounds when
	// fetching range_temp[index_temp].
	for i := range v { // 并没有在 for 作用域里面另外声明一个 v 的切片，range 是把 v 的值拷贝到内部的一个临时变量 tmp，所以下面append操作的依然是同一个 v
		fmt.Printf("%p, %d, %d\n", v, len(v), cap(v))
		v = append(v, i)
	}

	fmt.Printf("%p, %d, %d\n", v, len(v), cap(v))
	fmt.Println(v)
}
