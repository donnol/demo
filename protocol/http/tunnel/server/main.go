package main

import (
	"log"
	"net/http"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	handler := func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%#v\n", req)
		if req.Method == http.MethodConnect {
			host := req.URL.Host
			log.Println(host)
			w.Write([]byte(req.URL.String()))
		} else {
			w.Write([]byte("hello"))
		}
	}
	http.HandleFunc("/", handler)
	if err := http.ListenAndServeTLS(":8820", "cert.pem", "key.pem", nil); err != nil {
		log.Fatal(err)
	}
}
