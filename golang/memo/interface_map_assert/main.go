package main

import (
	"fmt"
	"time"
)

func main() {
	var m = make(map[string]interface{})

	m["day"] = time.Now().Day()

	// 当不存在该键时
	if v, ok := m["not_day"].(int); ok {
		fmt.Println(v)
	} else {
		fmt.Println(ok, v) // false, 0
	}
	//
	if v, ok := m["day"].(int); ok {
		fmt.Println(v) // 30
	}
}
