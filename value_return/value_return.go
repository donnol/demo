package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := a()
	fmt.Printf("ar: %p, %+v\n", &err, err)

	err = b()
	fmt.Printf("br: %p, %+v\n", &err, err)

	err = c()
	fmt.Printf("cr: %p, %+v\n", &err, err)
}

func a() (err error) {
	fmt.Printf("a1: %p, %+v\n", &err, err)
	defer func() {
		fmt.Printf("a5: %p, %+v\n", &err, err)
		err = nil
	}()
	if f, err := os.Open("abc"); err != nil {
		fmt.Printf("a2: %p, %+v\n", &err, err)
		defer func() {
			fmt.Printf("a4: %p, %+v\n", &err, err)
		}()
		// if作用域内有err，如果返回时不指定err，函数就会在函数作用域内寻找err，就会找到两个，从而报错：
		// 	err is shadowed during return
		// return
		// 但是，只要指定了err，那么函数就会返回这个err，忽略if作用域外的err
		// if the function executes a return statement with no arguments, the current values of the result parameters are used as the returned values.
		// -- 只有在return语句没有参数时，named return才会使用；有参数的时候就不用了，如下：
		return err
	} else {
		fmt.Printf("a3: %+v\n", f)
	}

	return
}

func b() (err error) {
	fmt.Printf("b1: %p, %+v\n", &err, err)
	f, err := os.Open("abc")
	fmt.Printf("b2: %p, %+v\n", &err, err)
	if err != nil {
		return
	}
	fmt.Printf("b3: %+v\n", f)
	return
}

func c() (err error) {
	fmt.Printf("c2: %p, %+v\n", &err, err)
	var err1 error
	fmt.Printf("c1: %p, %+v\n", &err1, err1)
	err = errors.New("abc")
	fmt.Printf("c2: %p, %+v\n", &err, err)

	// return err1
	return
}
