// 多维切片
package main

import "log"

func main() {
	// 一维
	mds1 := []string{}
	log.Printf("%+v\n", mds1)

	// 二维
	mds2 := [][]string{}
	log.Printf("%+v\n", mds2)

	mds2 = make([][]string, 10, 20)
	log.Printf("%+v\n", mds2)

	type column []string
	mdss := []column{}
	log.Printf("%+v\n", mdss)

	// TODO
}
