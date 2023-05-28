package main

import (
	"context"
	"fmt"
)

type FilterFunc func(ctx context.Context, req interface{}) (interface{}, error)

type HandlerFunc func(impl interface{}, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error)

type Method struct {
	Name string
	Func HandlerFunc
}

type Service1 interface {
	DoSomething(ctx context.Context, req string) (string, error)
}

type Service2 interface {
	DoAnotherThing(ctx context.Context, req int) (int, error)
}

type MyService1 struct{}

func (s *MyService1) DoSomething(ctx context.Context, req string) (string, error) {
	return "Hello, " + req, nil
}

type MyService2 struct{}

func (s *MyService2) DoAnotherThing(ctx context.Context, req int) (int, error) {
	return req * 2, nil
}

type HandlerImpl struct {
	Handler HandlerFunc
	Impl    interface{}
}

type Server struct {
	methods map[string]HandlerImpl
}

func NewServer() *Server {
	return &Server{
		methods: make(map[string]HandlerImpl),
	}
}

func (s *Server) RegisterMethod(method Method, serviceImpl interface{}) {
	s.methods[method.Name] = HandlerImpl{Handler: method.Func, Impl: serviceImpl}
}

func (s *Server) CallMethod(methodName string, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error) {
	method, ok := s.methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}
	return method.Handler(method.Impl, ctx, preFilter, postFilter)
}

func main() {
	server := NewServer()

	// service1实现并调用
	serviceImpl1 := &MyService1{}
	method1 := Method{
		Name: "DoSomething",
		Func: func(impl interface{}, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error) {
			service, ok := impl.(Service1)
			if !ok {
				return nil, fmt.Errorf("invalid service implementation")
			}
			req, err := preFilter(ctx, nil)
			if err != nil {
				return nil, err
			}
			resp, err := service.DoSomething(ctx, req.(string))
			if err != nil {
				return nil, err
			}
			return postFilter(ctx, resp)
		},
	}

	server.RegisterMethod(method1, serviceImpl1)

	// service2实现并调用
	serviceImpl2 := &MyService2{}
	method2 := Method{
		Name: "DoAnotherThing",
		Func: func(impl interface{}, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error) {
			service, ok := impl.(Service2)
			if !ok {
				return nil, fmt.Errorf("invalid service implementation")
			}
			req, err := preFilter(ctx, nil)
			if err != nil {
				return nil, err
			}
			resp, err := service.DoAnotherThing(ctx, req.(int))
			if err != nil {
				return nil, err
			}
			return postFilter(ctx, resp)
		},
	}

	server.RegisterMethod(method2, serviceImpl2)

	preFilter := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "World", nil
	}

	postFilter := func(ctx context.Context, resp interface{}) (interface{}, error) {
		return resp, nil
	}

	ctx := context.Background()
	response1, err := server.CallMethod("DoSomething", ctx, preFilter, postFilter)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response1:", response1)
	}

	preFilter2 := func(ctx context.Context, req interface{}) (interface{}, error) {
		return 5, nil
	}

	response2, err := server.CallMethod("DoAnotherThing", ctx, preFilter2, postFilter)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response2:", response2)
	}
}

/*
这里在第一个版本的基础上: 期望添加多个servie和filter进行业务处理，这也符合通常添加service的办法
两个不同的接口 Service1 和 Service2，分别包含 DoSomething 和 DoAnotherThing 方法。我们为每个服务注册了一个处理函数，每个处理函数可以处理各自服务的特定方法。
在这个示例中，MyService1 实现了 Service1 接口，而 MyService2 实现了 Service2 接口。
我们在 main 函数中分别调用了这两个服务的方法，并使用了不同的前置过滤器 preFilter 和 preFilter2 来处理不同类型的输入。后置过滤器 postFilter 保持通用，不对响应进行任何处理。
*/
