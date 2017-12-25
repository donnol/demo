package main

import (
	"net/http"
)

func main() {
	// 注册路由
	mux := NewMux()                // 新建 mux ，框架提供
	mux.Register(pattern, handler) // pattern 和 handler 由用户创建

	// 启动服务
	http.ListenAndServe(":9920", mux) // 标准库调用
}
