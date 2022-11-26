package main

import "fmt"

type Chain []Filter

type Filter func(next HandleFunc)

type HandleFunc func()

func (c Chain) Handle(handler HandleFunc) {
	for i := len(c) - 1; i >= 0; i-- {
		fmt.Println(i)
		curFilter, curHandler := c[i], handler
		handler = func() {
			curFilter(curHandler)
		}
	}
	handler()
}

func filter1(next HandleFunc) {
	fmt.Println("start handle in filter1")
	next()
	fmt.Println("end handle in filter1")
}

func filter2(next HandleFunc) {
	fmt.Println("start handle in filter2")
	next()
	fmt.Println("end handle in filter2")
}

func filter3(next HandleFunc) {
	fmt.Println("start handle in filter3")
	next()
	fmt.Println("end handle in filter3")
}

func handle() {
	fmt.Println("OK!")
}

/*
1. 构造链表，将一系列filter通过匿名函数的办法封装为和调用函数HandleFunc一致的签名类型，中间用next连接
2. 递归调用，从第一个filter从前到后调用至HandleFunc，然后返回进行逆序尾处理
*/
func main() {
	chain := Chain{filter1, filter2, filter3}
	chain.Handle(handle)
}
