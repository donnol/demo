package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"
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

	j := jsonLog{
		Msg:   "Hello, I am jd",
		Level: Info,
		Time:  time.Now(),
	}
	data, err = json.Marshal(j)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}

type jsonLog struct {
	Msg   string    `json:"msg"`
	Level Level     `json:"level"` // 该字段类型实现了json.Unmarshaler
	Time  time.Time `json:"time"`
	Extra string    `json:"extra,omitempty"` // 值为空时，忽略该字段
}

type Level uint32

const (
	Warning Level = 0
	Info    Level = 1
	Debug   Level = 2
)

// MarshalJSON implements json.Marshaler.MarashalJSON.
func (lv Level) MarshalJSON() ([]byte, error) {
	switch lv {
	case Warning:
		return []byte(`"warning: "`), nil
	case Info:
		return []byte(`"info: "`), nil
	case Debug:
		return []byte(`"debug: "`), nil
	default:
		return nil, fmt.Errorf("unknown level %v", lv)
	}
}

// UnmarshalJSON implements json.Unmarshaler.UnmarshalJSON.  It can unmarshal
// from both string names and integers.
func (lv *Level) UnmarshalJSON(b []byte) error {
	switch s := string(b); s {
	case "0", `"warning"`:
		*lv = Warning
	case "1", `"info"`:
		*lv = Info
	case "2", `"debug"`:
		*lv = Debug
	default:
		return fmt.Errorf("unknown level %q", s)
	}
	return nil
}
