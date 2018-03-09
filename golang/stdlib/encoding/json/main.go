package main

import (
	"bytes"
	"encoding/json"
	"log"
)

func main() {
	data := []byte(`{"key": "value"}`)
	valid := json.Valid(data)
	log.Printf("%v\n", valid)

	var m = make(map[string]string)
	err := json.Unmarshal(data, &m)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", m)

	dst := new(bytes.Buffer)
	err = json.Compact(dst, data) // 紧凑
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v, %s\n", dst, data) // {"key":"value"}, {"key": "value"} 中间的空格没了

	dst1 := new(bytes.Buffer)
	data1 := []byte(`{"key": "<a href='127.0.0.1'>click</a>"}`)
	json.HTMLEscape(dst1, data1)
	log.Printf("%v, %s\n", dst1, data1) // {"key": "\u003ca href="127.0.0.1"\u003eclick\u003c/a\u003e"}

	m1 := make(map[string]string)
	err = json.Unmarshal(dst1.Bytes(), &m1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", m1) // map[key:<a href='127.0.0.1'>click</a>]
}
