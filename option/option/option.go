// Package option 提供方
// 这种用法可以在确保 NewServer 不变的同时，保证扩展性
package option

import (
	"time"
)

// Server Server 类型
type Server struct {
	addr    string
	timeout time.Duration // 新加字段
}

// type Server http.Server

// Option 初始化 Server 的选项
type Option func(*Server)

// NewServer 新建 Server
func NewServer(option ...Option) *Server {
	server := new(Server)
	for _, opt := range option {
		opt(server)
	}
	return server
}

// SetServerAddr 设置 addr 字段值
func SetServerAddr(addr string) Option {
	return func(s *Server) {
		s.addr = addr
	}
}

// SetServerTimeout 新加字段 timeout 后，提供设置方法
func SetServerTimeout(d time.Duration) Option {
	return func(s *Server) {
		s.timeout = d
	}
}

// SetTimeout 修改 timeout
func (s *Server) SetTimeout(d time.Duration) {
	s.timeout = d
}
