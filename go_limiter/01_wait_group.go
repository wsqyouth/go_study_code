package main

import (
	"context"
	"fmt"

	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

func main() {
	g, ctx := errgroup.WithContext(context.Background())
	limiter := rate.NewLimiter(1, 1) // 速率和令牌数都设置为1

	for i := 0; i < 10; i++ {
		num := i
		g.Go(func() error {
			err := limiter.Wait(ctx) // 在每个 goroutine 内部等待令牌
			if err != nil {
				return err
			}
			fmt.Println(num)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	}
}

/*
思考: 看其他人的代码将limiter.Wait放到g.Go外面，感觉不大对劲，问了下chatgpt，它的缺点是:
它在启动新的 goroutine 之前调用了 limiter.Wait(ctx)，而不是在新的 goroutine 内部调用。
这可能会导致速率限制器的效果不如预期，因为 limiter.Wait(ctx) 只会在启动新的 goroutine 时阻塞，而不会在新的 goroutine 运行时阻塞。

改进版本:
将 limiter.Wait(ctx) 移动到了新的 goroutine 内部
创建了一个速率为1的速率限制器，并启动了10个 goroutine 来打印数字。
每个 goroutine 在运行时都会等待令牌桶中有可用的令牌，从而确保每秒只打印一个数字。
尽管有多个 goroutine 同时运行，但是由于速率限制器的作用，每秒只会打印一个数字。
*/
