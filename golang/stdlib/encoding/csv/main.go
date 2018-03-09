package main

import (
	"bytes"
	"encoding/csv"
	"log"
)

func main() {
	var w = new(bytes.Buffer)
	writer := csv.NewWriter(w)
	writer.Comma = ';' // 设置分隔符
	for _, record := range [][]string{
		[]string{"a", "b", "c"},
		[]string{"d", "e", "f"},
	} {
		err := writer.Write(record)
		if err != nil {
			log.Fatal(err)
		}
	}
	writer.Flush()

	log.Printf("%v\n", w.String())

	// var r = new(bytes.Buffer)
	reader := csv.NewReader(w)
	reader.Comma = ';' // 设置分隔符
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", records)
}
