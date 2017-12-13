package main

import (
	"github.com/coreos/etcd/etcdserver"
)

func main() {
	server, err := etcdserver.NewServer(etcdserver.ServerConfig{
		Name:    "a",
		DataDir: "tmp",
	})
	if err != nil {
		panic(err)
	}
	server.Start()
}
