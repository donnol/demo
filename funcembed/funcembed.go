package main

func main() {

}

type M struct{}

// 函数D的参数是f，f的类型是func(*M) func() (int, error)，这是一个函数类型，接受M的指针类型参数，返回func() (int, error)函数类型值
func D(f func(*M) func() (int, error)) error {
	var err error
	return err
}
