package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"reflect"
)

// M M
type M struct {
	Name string
}

func main() {
	var w = new(bytes.Buffer)
	encoder := gob.NewEncoder(w)
	e := M{"jd"}
	// err := encoder.Encode(e)
	err := encoder.EncodeValue(reflect.ValueOf(e))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v, %v\n", w.String(), w.Bytes())

	decoder := gob.NewDecoder(w)
	e1 := M{}
	// err = decoder.Decode(&e1)
	err = decoder.DecodeValue(reflect.ValueOf(&e1))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v\n", e1)
}
