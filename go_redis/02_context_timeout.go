package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id uint32) {
	select {
	case <-time.After(3 * time.Second):
		fmt.Printf("worker id: %v finished\n", id)
	case <-ctx.Done():
		fmt.Printf("worker time out. id:%v\n", id)
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		// 创建新的上下文
		cloneCtx, _ := context.WithTimeout(ctx, time.Second*5)
		// 使用 context.Background() 作为父上下文来创建 cloneCtx
		// cloneCtx, _ := context.WithTimeout(context.Background(), time.Second*5)
		go worker(cloneCtx, uint32(i))
	}
	time.Sleep(6 * time.Second)
	fmt.Println("hello world")
}

// 场景1: 所有goroutin使用同一个ctx, 当ctx超时后所有goroutin可以收到取消信号,打印time out消息
func sameContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		go worker(ctx, uint32(i))
	}
	time.Sleep(time.Second * 10)
}

// 场景2： 每个goroutine使用clone ctx, 这个 cloneCtx 的超时时间比原始的 ctx 长。因此，即使原始的 ctx 超时，这些 goroutine 也不会收到取消信号，而是会继续运行直到完成。
func diffContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		// 创建新的上下文
		cloneCtx, _ := context.WithTimeout(ctx, time.Second*5)
		// 使用 context.Background() 作为父上下文来创建 cloneCtx
		// cloneCtx, _ := context.WithTimeout(context.Background(), time.Second*5)
		go worker(cloneCtx, uint32(i))
	}
	time.Sleep(6 * time.Second)
}

/*
这里主要考虑的问题是另外起一个协程删除分布式锁时，是否需要clone ctx
每个 worker goroutine 在开始时都会创建自己的 cloneCtx，这个 cloneCtx 的超时时间比原始的 ctx 长。因此，即使原始的 ctx 超时，这些 goroutine 也不会收到取消信号，而是会继续运行直到完成。

假设你的操作在5秒内没有完成，上下文被取消了。如果你没有创建新的上下文，那么在 defer 中的 redisClient.Del(ctx, key) 也会收到取消信号，可能无法成功删除锁。这就是为什么你需要创建新的上下文，以确保可以成功删除锁。

然而这种做法也有一个潜在的问题。假设你的程序在主 goroutine 中等待所有的子 goroutine 完成。通常，你可以通过取消上下文来通知所有的子 goroutine 尽快结束。
但是，如果你在子 goroutine 中创建了新的上下文，那么这个子 goroutine 就不会收到取消信号，可能会继续运行，即使主 goroutine 已经结束了。这可能会导致程序无法正确地结束，或者浪费资源。
通俗的说就是不受控制，可能有问题
*/
