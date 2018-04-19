package main

import (
	"io/ioutil"
	"log"
)

func main() {
	// discard
	// var Discard io.Writer = devNull(0) -- devNull(0) 类型转换(type devNull int)
	data := []byte("Hello, I am jd")
	_, err := ioutil.Discard.Write(data) // 所有写的操作都会成功，但是什么都不做
	if err != nil {
		log.Fatal(err)
	}
}
