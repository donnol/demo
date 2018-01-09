package main

import (
	"log"
	"net/rpc"

	"demo/rpc/server"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &server.Args{7, 8}
	var reply int64
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Println("result is", reply)

	var quo server.Quotient
	err = client.Call("Arith.Divide", args, &quo)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Println("result is", quo)

	quotient := new(server.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		log.Fatal("async aritch error:", replyCall.Error)
	}
	log.Printf("result is %v, %v\n", *replyCall, replyCall.Reply.(*server.Quotient))
}
