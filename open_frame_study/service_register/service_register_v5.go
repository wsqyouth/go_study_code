package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type ErrorMethodNotFound struct {
	MethodName string
}

func (e *ErrorMethodNotFound) Error() string {
	return fmt.Sprintf("method not found: %s", e.MethodName)
}

type ErrorInvalidMethodName struct {
	MethodName string
}

func (e *ErrorInvalidMethodName) Error() string {
	return fmt.Sprintf("invalid method name: %s", e.MethodName)
}

type ErrorInvalidRequest struct {
	Message string
}

func (e *ErrorInvalidRequest) Error() string {
	return fmt.Sprintf("invalid request: %s", e.Message)
}

type MiddlewareFunc func(ctx context.Context, req interface{}, next HandlerFunc) (interface{}, error)

type Service interface {
	RegisterMethods(server *Server)
}

type HandlerFunc func(ctx context.Context, req interface{}) (interface{}, error)

type Server struct {
	methods map[string]HandlerFunc
}

func NewServer() *Server {
	return &Server{
		methods: make(map[string]HandlerFunc),
	}
}

func (s *Server) RegisterService(service Service) {
	service.RegisterMethods(s)
}

func (s *Server) RegisterMethod(methodName string, handlerFunc HandlerFunc) {
	s.methods[methodName] = handlerFunc
}

func (s *Server) CallMethod(methodName string, ctx context.Context, req interface{}) (interface{}, error) {
	handlerFunc, ok := s.methods[methodName]
	if !ok {
		return nil, &ErrorMethodNotFound{MethodName: methodName}
	}
	return handlerFunc(ctx, req)
}

// UseMiddleware 比较巧妙,使用闭包同时将中间件都放入s.methods中
func (s *Server) UseMiddleware(middleware MiddlewareFunc) {
	wrappedMethods := make(map[string]HandlerFunc)
	for methodName, handlerFunc := range s.methods {
		wrappedMethods[methodName] = func(originalHandler HandlerFunc) HandlerFunc {
			return func(ctx context.Context, req interface{}) (interface{}, error) {
				return middleware(ctx, req, originalHandler)
			}
		}(handlerFunc) // Pass handlerFunc as an argument to the closure
	}
	s.methods = wrappedMethods
}

type MyService1 struct{}

func (s *MyService1) RegisterMethods(server *Server) {
	server.RegisterMethod("DoSomething", s.DoSomething)
}

func (s *MyService1) DoSomething(ctx context.Context, req interface{}) (interface{}, error) {
	str, ok := req.(string)
	if !ok {
		return nil, &ErrorInvalidRequest{Message: "request must be a string"}
	}
	if str == "" {
		return nil, &ErrorInvalidRequest{Message: "request cannot be empty"}
	}
	return "Hello, " + str, nil
}

type MyService2 struct{}

func (s *MyService2) RegisterMethods(server *Server) {
	server.RegisterMethod("DoAnotherThing", s.DoAnotherThing)
}

func (s *MyService2) DoAnotherThing(ctx context.Context, req interface{}) (interface{}, error) {
	num, ok := req.(int)
	if !ok {
		return nil, &ErrorInvalidRequest{Message: "request must be an integer"}
	}
	if num < 0 {
		return nil, &ErrorInvalidRequest{Message: "request cannot be negative"}
	}
	return num * 2, nil
}

func LoggingMiddleware(ctx context.Context, req interface{}, next HandlerFunc) (interface{}, error) {
	log.Printf("request: %v", req)

	startTime := time.Now()
	resp, err := next(ctx, req)
	endTime := time.Now()

	log.Printf("response: %v, error: %v, elapsed time: %v", resp, err, endTime.Sub(startTime))

	return resp, err
}

func AuthorizationMiddleware(ctx context.Context, req interface{}, next HandlerFunc) (interface{}, error) {
	// TODO: implement authorization logic
	return next(ctx, req)
}

func main() {
	server := NewServer()

	service1 := &MyService1{}
	service1.RegisterMethods(server)

	service2 := &MyService2{}
	service2.RegisterMethods(server)

	server.UseMiddleware(LoggingMiddleware)
	server.UseMiddleware(AuthorizationMiddleware)

	req1 := "world"
	resp1, err1 := server.CallMethod("DoSomething", context.Background(), req1)
	if err1 != nil {
		log.Printf("error: %v", err1)
	} else {
		log.Printf("response: %v", resp1)
	}

	req2 := 3
	resp2, err2 := server.CallMethod("DoAnotherThing", context.Background(), req2)
	if err2 != nil {
		log.Printf("error: %v", err2)
	} else {
		log.Printf("response: %v", resp2)
	}
}

/*
为了优化这段代码，我们可以采取以下几个步骤：

使用更具描述性的错误类型，以便于调试和错误处理。
在这段代码中，所有的错误都是使用 fmt.Errorf() 函数创建的通用错误。这样做虽然简单，但是不利于调试和错误处理。我们可以为不同的错误类型创建自定义错误类型，以便于更好地理解和处理错误。

例如，我们可以为方法未找到、无效的方法名称、无效的请求参数等错误创建自定义错误类型。

将服务和方法的注册与调用分离，以便于扩展和模块化。
在这段代码中，服务和方法的注册和调用是紧密耦合的。这样做虽然简单，但是不利于扩展和模块化。我们可以将服务和方法的注册和调用分离，以便于更好地扩展和组织代码。

例如，我们可以为服务创建一个接口，并为每个服务实现一个结构体。然后，我们可以将服务注册到服务器中，并为每个服务方法创建一个处理程序函数。最后，我们可以通过调用服务器的 CallMethod() 方法来调用服务方法。

使用中间件模式替换过滤器，以便于扩展和组合功能。
在这段代码中，过滤器是用于处理请求参数的函数列表。这样做虽然简单，但是不利于扩展和组合功能。我们可以使用中间件模式替换过滤器，以便于更好地扩展和组合功能。

例如，我们可以为每个中间件创建一个处理程序函数，并将它们链接在一起。然后，我们可以将中间件链传递给服务方法处理程序函数，以便于处理请求参数。最后，我们可以通过调用服务器的 CallMethod() 方法来调用服务方法。
*/
