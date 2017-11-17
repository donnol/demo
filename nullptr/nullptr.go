package main

func main() {
	run()
}

func run() {
	panic("null pointer")
	var i *int
	*i = 1
}
