package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	testCase := []jsonStruct{
		{"a", "b"},
		{"c", 1},
	}
	jsonStr, err := testJson(testCase)
	if err != nil {
		panic(err)
	}
	b := bytes.NewBuffer([]byte{})
	err = json.Indent(b, []byte(jsonStr), "", "	")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", b)
}

type jsonStruct struct {
	key   string
	value interface{}
}

func testJson(j []jsonStruct) (ret string, err error) {
	suf := ", "
	jlen := len(j)
	for i := 0; i < jlen; i++ {
		switch j[i].value.(type) {
		case string:
			ret += fmt.Sprintf("\"%s\": \"%s\"%s", j[i].key, j[i].value, suf)
		case int:
			ret += fmt.Sprintf("\"%s\": %d%s", j[i].key, j[i].value, suf)
		default:
			panic("error")
		}
	}
	ret = fmt.Sprintf("{%s}", strings.TrimRight(ret, suf))
	return
}
