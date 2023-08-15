package main

import (
	"fmt"
)

type Server struct {
	address string
	port    int
}

type ServerOption func(*Server)

func WithAddress(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

func WithPort(port int) ServerOption {
	return func(s *Server) {
		s.port = port
	}
}

func NewServer(opts ...ServerOption) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func main() {
	s := NewServer(WithAddress("localhost"), WithPort(8080))
	fmt.Println(s.address) // 输出: localhost
	fmt.Println(s.port)    // 输出: 8080
}

/*
opton设计模式：
解决问题：
1. 参数过多  2. 易于扩展 3. 增加可维护性
*/
