package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	host, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, " + host))
	})
	port := ":8520"

	log.Printf("host: %s\nlistening port: %s\n", host, port)
	err = http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err)
	}
}
