package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()
	limiter := rate.NewLimiter(10, 1)
	start := time.Now()
	//模拟高并发请求
	for i := 0; i < 100; i++ {
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Println("limiter err")
			return
		}
		fmt.Printf("sum: %v\n", Add(30, 20))
	}
	fmt.Println(time.Since(start))
}

func Add(a, b int64) int64 {
	return a + b
}
