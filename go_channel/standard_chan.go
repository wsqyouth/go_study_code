package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/time/rate"
)

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 方案1: 控制并发数量
	// rateLimitedProcessing1(items, 3)

	// 方案2: 通过rate+sync.WaitGroup控制并发处理速率
	limiter := rate.Limit(3)
	// rateLimitedProcessing2(items, limiter)

	// 方案3: 通过rate+errgroup控制并发处理速率
	maxGoroutine := runtime.NumGoroutine()
	rateLimitedProcessing3(items, limiter, maxGoroutine)
}

func processItem(item int) {
	fmt.Println("Processing item", item)
	time.Sleep(time.Second) // 模拟处理时间
}

func rateLimitedProcessing1(items []int, limit int) {
	var wg sync.WaitGroup
	ch := make(chan struct{}, limit)

	for _, item := range items {
		ch <- struct{}{}
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			processItem(item)
			<-ch
		}(item)
	}
	wg.Wait()
}

func rateLimitedProcessing2(items []int, r rate.Limit) {
	var wg sync.WaitGroup
	limiter := rate.NewLimiter(r, 3)

	for _, item := range items {
		wg.Add(1)
		go func(item int) {
			defer wg.Done()
			if err := limiter.Wait(context.Background()); err != nil {
				fmt.Println("Error:", err)
				return
			}
			processItem(item)
		}(item)
	}
	wg.Wait()
}

func rateLimitedProcessing3(items []int, r rate.Limit, maxGoroutine int) {
	g, ctx := errgroup.WithContext(context.Background())
	limiter := rate.NewLimiter(r, 1)
	ch := make(chan int)

	for i := 0; i < maxGoroutine; i++ {
		g.Go(func() error {
			for item := range ch {
				if err := limiter.Wait(ctx); err != nil {
					return err
				}
				processItem(item)
			}
			return nil
		})
	}

	for _, item := range items {
		ch <- item
	}
	close(ch)

	g.Wait()
}

/*
以上三种方法都可以实现对处理速度的限制，同时也能处理并发。第一种方法使用channel作为信号量来限制并发数，第二种和第三种方法使用errgroup来处理并发，并且使用rate.Limiter来限制处理速度。

使用channel作为信号量来限制并发数：
优点：这种方法简单直观，易于理解和实现。通过控制channel的容量，可以很容易地限制并发数。
缺点：这种方法只能限制并发数，不能直接限制处理速度。如果处理函数的执行时间不一致，可能会导致处理速度波动。此外，如果处理函数中有panic，可能会导致goroutine泄漏。

使用errgroup来处理并发，并且使用rate.Limiter来限制处理速度：
优点：这种方法可以同时限制并发数和处理速度，更加灵活。errgroup可以确保所有goroutine都正确地返回，避免goroutine泄漏。rate.Limiter可以更精确地控制处理速度。
缺点：这种方法需要更多的代码来实现，并且理解起来稍微复杂一些。如果处理函数的执行时间不一致，可能会导致处理速度波动。

使用channel和errgroup来处理并发，并且使用rate.Limiter来限制处理速度：
优点：这种方法结合了前两种方法的优点，既可以限制并发数，又可以限制处理速度。errgroup可以确保所有goroutine都正确地返回，避免goroutine泄漏。rate.Limiter可以更精确地控制处理速度。
缺点：这种方法需要最多的代码来实现，并且理解起来最复杂。如果处理函数的执行时间不一致，可能会导致处理速度波动。

在实际工程中，应该根据具体需求来选择合适的设计。如果只需要限制并发数，且处理函数的执行时间比较一致，可以选择第一种方法。如果需要更精确地控制处理速度，可以选择第二种或第三种方法。如果处理函数的执行时间不一致，或者有可能出现错误或panic，应该选择使用errgroup的
*/
