// 调用方
package main

import (
	"log"

	"./option"
)

func main() {
	// 公有字段可以在这里写，私有字段就不行了
	// setServerAddr := func(addr string) option.Option {
	// 	return func(s *option.Server) {
	// 		s.Addr = addr // 当 addr 是私有字段，会报错 s.Addr undefined (type *option.Server has no field or method Addr, but does have option.addr)
	// 	}
	// }

	// 当 addr 是私有字段，由包提供初始化方法
	setServerAddr := option.SetServerAddr

	server := option.NewServer(setServerAddr(":9520"))
	log.Printf("server: %+v\n", server)

	// 当需要增加 timeout 配置时，向提供方提需求，然后提供方给予 SetServerTimeout 函数
	setServerTimeout := option.SetServerTimeout

	server = option.NewServer(setServerTimeout(1))
	log.Printf("server: %+v\n", server)

	// 同时设置 addr 和 timeout
	server = option.NewServer(setServerAddr(":9920"), setServerTimeout(5))
	log.Printf("server: %+v\n", server)

	// 上面的方法确保 Server 被创建后，就无法被修改了

	// 但是，创建出来 Server 后，有人跟提供方说，我想修改它的 timeout，提供方给予 Server.SetTimeout 方法
	server.SetTimeout(10)
	log.Printf("server: %+v\n", server)
}
