package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	const numSize int = 20
	nums := makeRange(1, numSize)
	ch := make(chan int, numSize)

	ctx := context.Background()
	// 1. 主线程生成
	//generateNumCh(nums, ch)
	// 2. 多线程生产
	batchGenerateNumCh(nums, ch)
	close(ch) //不再填充后要close，否则读线程一直读取不到数据会deadlock
	fmt.Println("len:", len(ch))
	time.Sleep(time.Second)

	// 1. 使用一个线程消费
	//go print(ch)
	// 2.使用多个线程消费,errgroup
	if err := batchCousumeWithErrGroup(ctx, ch); err != nil {
		panic("err")
	}
}

// makeRange 生成1-10000的数组
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// 方案1: 数据顺序放到channel里
func generateNumCh(in []int, ch chan int) {
	for _, each := range in {
		ch <- each
	}
}

// 方案2：多线程生产
func batchGenerateNumCh(nums []int, ch chan int) {
	var wg sync.WaitGroup
	for _, num := range nums {
		tmpNum := num //// 这个在loop内是需要的，否则Loop variable captured by func literal
		wg.Add(1)
		go func() {
			ch <- tmpNum
			wg.Done()
		}()
	}
	wg.Wait()
}

// 方案1： 单线程消费
func print(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}

// 方案2: 使用errgroup进行并发消费
func batchCousumeWithErrGroup(ctx context.Context, ch chan int) error {
	maxGoroutine := 4
	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i < maxGoroutine; i++ {
		g.Go(func() (e error) {
			for val := range ch {
				fmt.Println(val)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
