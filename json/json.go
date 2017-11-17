package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	m := M{"jd", 28, "123", "12345678901"}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	var m2 M
	err = json.Unmarshal([]byte(`{"name":"jd","Age":28,"phone":"12345678901"}`), &m2)
	if err != nil {
		panic(err)
	}
	fmt.Println(m2)
}

type M struct {
	Name  string
	Age   int `json:"age"`
	pwd   string
	Phone string `json:"Phone"`
}
