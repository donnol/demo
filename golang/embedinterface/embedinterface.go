package main

func main() {

}

// 一个接口
type I interface {
	Talk(s string) error
}

// 当已经有很多结构体实现了该接口，如果想扩展该接口，添加一个Listen()方法，则可以使用
type EI interface {
	I
	Listen() (string, error)
}
