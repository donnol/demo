package main

import "fmt"

var sptr *string

type PM struct {
	// pm *pm
	*pm
}

type pm struct {
	Title *string
}

func main() {
	// 	panic: runtime error: invalid memory address or nil pointer dereference
	// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x4831b0]
	// 1
	// fmt.Println(*sptr)
	// 2
	// instance := PM{}
	// fmt.Println(instance.Title)
	// fmt.Println(instance.pm.Title)

	instance := PM{pm: &pm{}}
	fmt.Println(instance.Title)
	fmt.Println(instance.pm.Title)

	var a, b = "a", "b"
	_ = a
	_ = b

	instance.Title = &a
	fmt.Println(instance.Title, *instance.Title)
	instance.pm.Title = &b
	fmt.Println(instance.pm.Title, *instance.pm.Title)
}
