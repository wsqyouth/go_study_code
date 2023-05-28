package main

import (
	"context"
	"fmt"
	"reflect"
)

type FilterFunc func(ctx context.Context, req interface{}) (interface{}, error)

type HandlerFunc func(impl interface{}, methodName string, ctx context.Context, filters []FilterFunc) (interface{}, error)

type Method struct {
	Name string
	Func HandlerFunc
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

func (s *Server) RegisterMethod(methodName string, serviceImpl interface{}, handlerFunc HandlerFunc) {
	s.methods[methodName] = HandlerImpl{Handler: handlerFunc, Impl: serviceImpl}
}

func (s *Server) CallMethod(methodName string, ctx context.Context, filters []FilterFunc) (interface{}, error) {
	method, ok := s.methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}
	return method.Handler(method.Impl, methodName, ctx, filters)
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

func genericHandler(impl interface{}, methodName string, ctx context.Context, filters []FilterFunc) (interface{}, error) {
	req := interface{}(nil)
	var err error

	for _, filter := range filters {
		req, err = filter(ctx, req)
		if err != nil {
			return nil, err
		}
	}

	method := reflect.ValueOf(impl).MethodByName(methodName)
	if !method.IsValid() {
		return nil, fmt.Errorf("invalid method name: %s", methodName)
	}

	inputs := []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(req)}
	outputs := method.Call(inputs)

	err, _ = outputs[1].Interface().(error)
	if err != nil {
		return nil, err
	}

	resp := outputs[0].Interface()
	return resp, nil
}

func main() {
	server := NewServer()

	serviceImpl1 := &MyService1{}
	server.RegisterMethod("DoSomething", serviceImpl1, genericHandler)

	serviceImpl2 := &MyService2{}
	server.RegisterMethod("DoAnotherThing", serviceImpl2, genericHandler)

	preFilter := func(ctx context.Context, req interface{}) (interface{}, error) {
		return "World", nil
	}

	postFilter := func(ctx context.Context, resp interface{}) (interface{}, error) {
		return resp, nil
	}

	ctx := context.Background()
	response1, err := server.CallMethod("DoSomething", ctx, []FilterFunc{preFilter, postFilter})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response1:", response1)
	}

	preFilter2 := func(ctx context.Context, req interface{}) (interface{}, error) {
		return 5, nil
	}

	response2, err := server.CallMethod("DoAnotherThing", ctx, []FilterFunc{preFilter2, postFilter})
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Response2:", response2)
	}
}

/*
想法: 之前的代码需要明确定义前后过滤器，要求使用链式处理继续优化代码
提问：如果你是一名golang高级工程师，请问最新的代码还有哪些地方要优化的，请列举出几点并优化
使用更具描述性的错误类型，以便于调试和错误处理。
将服务和方法的注册与调用分离，以便于扩展和模块化。
使用中间件模式替换过滤器，以便于扩展和组合功能。
*/
