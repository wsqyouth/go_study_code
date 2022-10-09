package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	const numSize int = 20
	nums := makeRange(1, numSize)
	ch := make(chan int, numSize)

	ctx := context.Background()
	generateNumCh(nums, ch)
	fmt.Println(len(ch))
	time.Sleep(time.Second)
	//go print(ch)
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
	close(ch)
	if err := g.Wait(); err != nil {
		panic("errgroup err")
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

// 数据顺序放到channel里
func generateNumCh(in []int, ch chan int) {
	for _, each := range in {
		ch <- each
	}
}

func print(ch chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
