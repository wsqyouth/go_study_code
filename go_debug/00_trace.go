package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

//trace 编译
func main() {
	//1. 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()
	// 2. 启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("hello gmp")
	//3. 停止trace
	trace.Stop()
}

// 测试:通过go tool trace 打开trace文件进行分析
//go run trace.go
//go tool trace trace.out,打开本地浏览器分析
// 可以分析出G0,M0这些的时间关系
