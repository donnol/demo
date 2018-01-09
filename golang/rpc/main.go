package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"demo/rpc/server"
)

func main() {
	arith := new(server.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	err = http.Serve(l, nil)
	if err != nil {
		log.Fatal("serve err is", err)
	}
}
