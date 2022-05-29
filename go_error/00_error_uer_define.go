package main

import (
	"fmt"
	"runtime"
)

func main() {
	content, err := openFile()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}
}

// 模拟错误
func openFile() ([]byte, error) {
	return nil, New("file not exist")
}

// 自定义错误类型，实现error接口
func New(msg string) error {
	return &errorString{
		msg:   msg,
		stack: callers(),
	}
}

type stack []uintptr

func callers() *stack {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(3, pcs[:])
	var st stack = pcs[0:n]
	return &st
}

type errorString struct {
	msg   string
	stack *stack
}

func (e *errorString) Error() string {
	return e.msg
}

//ref: https://www.flysnow.org/2019/01/01/golang-error-handle-suggestion.html
