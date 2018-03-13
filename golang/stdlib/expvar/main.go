package main

import (
	"expvar"
	"log"
	"net/http"
)

func main() {
	intVar := expvar.NewInt("MyInt")
	intVar.Add(1)
	log.Printf("%#v, %d\n", intVar, intVar.Value())

	intVar.Set(2)
	log.Printf("%#v, %d\n", intVar, intVar.Value())

	f := func(kv expvar.KeyValue) {
		log.Printf("%+v\n", kv)
	}
	expvar.Do(f)

	v := expvar.Get("MyInt")
	expvar.Publish("pub", v)

	handler := expvar.Handler()
	err := http.ListenAndServe(":8822", handler)
	if err != nil {
		log.Fatal(err)
	}
}
