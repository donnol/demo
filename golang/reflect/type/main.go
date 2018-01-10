package main

import (
	"reflect"
	"util"
)

func main() {
	// 管道
	recvDir := reflect.RecvDir
	sendDir := reflect.SendDir
	bothDir := reflect.BothDir

	// 字符串
	str := "hello"
	strType := reflect.TypeOf(str)

	util.Print(recvDir, sendDir, bothDir, strType)
}
