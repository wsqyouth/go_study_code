package main

import "fmt"

type HandlerFunc func(*Request)

type Request struct {
	url      string
	handlers []HandlerFunc
	index    int // 新增
}

func (r *Request) Use(handlerFunc HandlerFunc) {
	r.handlers = append(r.handlers, handlerFunc)
}

// 新增
func (r *Request) Next() {
	r.index++
	for r.index < len(r.handlers) {
		r.handlers[r.index](r)
		r.index++
	}
}

// 修改
func (r *Request) Run() {
	r.index = -1
	r.Next()
}

// 测试
// 输出 1 2 3 11
func main() {
	r := &Request{}
	//r.Use(func(r *Request) {
	//	fmt.Print(1, " ")
	//	r.Next()
	//	fmt.Print(11, " ")
	//})
	//r.Use(func(r *Request) {
	//	fmt.Print(2, " ")
	//})
	r.Use(func(r *Request) {
		fmt.Print(3, " ")
	})
	r.Run()
}

/*
这个是gin框架中使用职责链模式的核心实现
*/
