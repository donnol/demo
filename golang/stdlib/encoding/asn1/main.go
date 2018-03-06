package main

import (
	"encoding/asn1"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	type Model struct {
		Name string
	}

	for _, name := range []string{
		"jd",    // [48 4 19 2 106 100]
		"jdlau", // [48 7 19 5 106 100 108 97 117] 5 表示 5 字节长度
		// 其中 19 表示 TagPrintableString；后面紧接的数字 表示 内容的字节长度
	} {
		model := Model{
			Name: name,
		}
		data, err := asn1.Marshal(model)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(data, string(data))

		var model2 Model
		rest, err := asn1.Unmarshal(data, &model2)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(model2, rest, string(rest))

		var rawValue asn1.RawValue
		rest, err = asn1.Unmarshal(data, &rawValue)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%+v, %v, %s", rawValue, rest, rest)
	}
	// 真正吸引人的是自由和满足
	// 自由是对外界的感受
	// 满足是对自我的实现
}
