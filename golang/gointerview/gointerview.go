package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	// for _, stu := range stus {
	// 	m[stu.Name] = &stu
	// }
	for i, _ := range stus {
		stu := stus[i]
		m[stu.Name] = &stu

		fmt.Printf("%p, %p\n", &stus[i], &stu)
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("key=%s,value=%v \n", k, v)
	}
}
