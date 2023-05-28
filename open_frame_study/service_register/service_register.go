package main

import (
	"context"
	"fmt"
)

type FilterFunc func(ctx context.Context, req interface{}) (interface{}, error)

type HandlerFunc func(impl interface{}, ctx context.Context, f FilterFunc) (interface{}, error)

type Method struct {
	Name string
	Func HandlerFunc
}

type Service interface {
	DoSomething(ctx context.Context, req string) (string, error)
}

type MyService struct{}

func (s *MyService) DoSomething(ctx context.Context, req string) (string, error) {
	return "Hello, " + req, nil
}

type Server struct {
	handlers map[string]HandlerFunc
	impls    map[string]interface{}
}

func NewServer() *Server {
	return &Server{
		handlers: make(map[string]HandlerFunc),
		impls:    make(map[string]interface{}),
	}
}

func (s *Server) RegisterMethod(method Method, serviceImpl interface{}) {
	h := method.Func
	s.handlers[method.Name] = h
	s.impls[method.Name] = serviceImpl
}

func (s *Server) CallMethod(methodName string, ctx context.Context, f FilterFunc) (interface{}, error) {
	handler, ok := s.handlers[methodName]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}
	impl, ok := s.impls[methodName]
	if !ok {
		return nil, fmt.Errorf("implementation not found: %s", methodName)
	}
	return handler(impl, ctx, f)
}

func main() {
	server := NewServer()

	serviceImpl := &MyService{}
	method := Method{
		Name: "DoSomething",
		Func: func(impl interface{}, ctx context.Context, f FilterFunc) (interface{}, error) {
			service, ok := impl.(Service)
			if !ok {
				return nil, fmt.Errorf("invalid service implementation")
			}
			req, err := f(ctx, nil)
			if err != nil {
				return nil, err
			}
			return service.DoSomething(ctx, req.(string))
		},
	}

	server.RegisterMethod(method, serviceImpl)

	filter := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "World", nil
	}

	ctx := context.Background()
	response, err := server.CallMethod("DoSomething", ctx, filter)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response:", response)
	}
}

/*
这里借助gpt实现一个简单的service注册框架代码:
1. 思想：依赖注入
2. 添加了业务处理函数和filter，这里可以继续扩展,多个sercive及filter
3. service必须实现interface的方法
*/
