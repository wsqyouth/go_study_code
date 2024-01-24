package main

import (
	"context"
	"fmt"
	"log"
	"runtime/debug"
	"time"
)

func main() {
	// 当panicDemo函数触发panic时，Go函数会捕获到panic并使用log.Printf记录错误信息。
	Go(context.Background(), panicDemo)
	time.Sleep(time.Second * 10)
	fmt.Println("hello")
}

func panicDemo() {
	var x, y = 1, 0
	defer fmt.Println("exits for panicking:", x)
	x = x / y // will panic
	x++       // unreachable
}

// Go calls the given function in a new goroutine. panic is recovered.
func Go(ctx context.Context, f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic: %v, stacktrace: %s", r, string(debug.Stack()))
			}
		}()

		f()
	}()
}

/*
使用原始关键字go开启goroutine可能会有panic，因此需要对其封装，后续可以使用这个进行保护

参考资料:
[Defer, Panic, and Recover] https://go.dev/blog/defer-panic-and-recover
[Go101 好资料] https://go101.org/article/panic-and-recover-more.html
*/
