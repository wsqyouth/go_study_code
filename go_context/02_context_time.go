package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go slowOperation(ctx)
	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("goroutine:", runtime.NumGoroutine())
		}
	}()
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	time.Sleep(4 * time.Second)
	fmt.Println("Hello World")
}

func slowOperation(ctx context.Context) {
	done := make(chan int, 1)
	go func() {
		dur := time.Duration(rand.Intn(5)+1) * time.Second
		time.Sleep(dur)
		done <- 1
	}()
	select {
	case <-ctx.Done():
		fmt.Println("slowOperation timeout: ", ctx.Err)
	case <-done:
		fmt.Println("complete work")
	}
}

//ref: https://go.cyub.vip/concurrency/context.html
