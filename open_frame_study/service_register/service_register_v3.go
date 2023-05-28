package main

import (
	"context"
	"fmt"
	"reflect"
)

type FilterFunc func(ctx context.Context, req interface{}) (interface{}, error)

type HandlerFunc func(impl interface{}, methodName string, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error)

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

func (s *Server) CallMethod(methodName string, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error) {
	method, ok := s.methods[methodName]
	if !ok {
		return nil, fmt.Errorf("method not found: %s", methodName)
	}
	return method.Handler(method.Impl, methodName, ctx, preFilter, postFilter)
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

/*
它使用反射来调用服务实现的方法。核心还是map,只是通过反射将其设计更通用
这样，我们可以为每个服务定义一个独立的接口，并在注册时提供接口的方法名。
这使得代码更具可扩展性，可以根据实际项目需求添加更多的服务和方法。同时，您可以根据需要为每个服务定制前后过滤器。
*/
func genericHandler(impl interface{}, methodName string, ctx context.Context, preFilter, postFilter FilterFunc) (interface{}, error) {
	req, err := preFilter(ctx, nil)
	if err != nil {
		return nil, err
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
	return postFilter(ctx, resp)
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
想法：这里继续扩展，除了让这个程序能够扩展不同的service和filter,同时要求提供一个通用的机制调用服务实现方法
然后chatgpt给出了这个设计，总体还是不错，下一步优化filter
*/
