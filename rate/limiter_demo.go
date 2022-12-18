package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*10))
	defer cancel()
	// 默认设置1秒钟1个请求
	obtainTokenPerSecond := 1
	limiter := rate.NewLimiter(rate.Limit(obtainTokenPerSecond), int(obtainTokenPerSecond))
	maxGoroutine := 5
	g, ctx := errgroup.WithContext(ctx)
	start := time.Now()
	for i := 0; i < int(maxGoroutine); i++ {
		g.Go(func() (e error) {
			if err := processHighReq(ctx, limiter); err != nil {
				fmt.Println("processHighReq err")
				return err
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return
	}
	fmt.Println(time.Since(start))
}

// 模拟高并发请求
func processHighReq(ctx context.Context, limiter *rate.Limiter) error {
	for i := 0; i < 10; i++ {
		err := limiter.Wait(ctx)
		if err != nil {
			fmt.Println("limiter err")
			select {
			case <-ctx.Done(): // 时间结束了
				return ctx.Err()
			default:
				// 超频了，等一等
				fmt.Println("overflow limiter")
				time.Sleep(time.Second)
			}
		}
		fmt.Printf("sum: %v\n", Add(30, 20))
	}
	return nil
}

func Add(a, b int64) int64 {
	return a + b
}
