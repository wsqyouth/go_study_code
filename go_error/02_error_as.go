package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *MyError) Error() string {
	return e.Msg
}

func (e *MyError) ErrorCodeEqual(err error) bool {
	var myErr *MyError

	ok := errors.As(err, &myErr)
	if !ok {
		return false
	}
	return myErr.Code == e.Code
}

func main() {
	var err = &MyError{Code: 10600, Msg: "my error type"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)

	refErr := &MyError{Code: 10600}
	if refErr.ErrorCodeEqual(err2) {
		println("MyError is on the chain of err2 ")
		return
	}

	println("MyError is not on the chain of err2 ")
}

/*
总结：这里看了老白的书，核心就是这一句:
如果error类型变量的底层错误值是一个包装错误，那么errors.As方法会沿着该包装错误所在错误链与链上所有被包装的错误的类型进行比较，直至找到一个匹配的错误类型
*/
