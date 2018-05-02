package main

import "log"

func main() {
	m := &M{
		m{
			"jd",
			18,
			"male",
		},
		"ja",
	}
	log.Println(m)

	m2 := &M{}
	m2.Name = "jc"
	log.Println(m2)
}

// M M 类型
type M struct {
	m
	Name string
}

type m struct {
	Name string
	Age  int
	Sex  string
}
