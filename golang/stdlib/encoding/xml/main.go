package main

import (
	"encoding/xml"
	"log"
)

// M M
type M struct {
	// name string // <M></M>
	Name string
}

func main() {
	m := M{"jd"}
	data, err := xml.Marshal(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data) // <M><Name>jd</Name></M>

	var m1 M
	err = xml.Unmarshal(data, &m1)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", m1)
}
