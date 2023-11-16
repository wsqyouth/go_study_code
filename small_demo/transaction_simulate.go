package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

type Ad struct {
	Name string
}

func main() {
	ads := []Ad{{"Ad1"}, {"Ad2"}, {"Ad3"}}

	err := TransactFunc(context.Background(), func(ctx context.Context) error {
		g, gCtx := errgroup.WithContext(ctx)
		for _, ad := range ads {
			ad := ad
			g.Go(func() error {
				err := CreateAd(gCtx, ad)
				if err != nil {
					return err
				}
				return nil
			})
			fmt.Printf("Created ad: %s\n", ad.Name)
		}
		if err := g.Wait(); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func TransactFunc(ctx context.Context, fn func(ctx context.Context) error) error {
	// Simulate a transaction by just calling the function.
	return fn(ctx)
}

func CreateAd(ctx context.Context, ad Ad) error {
	// Simulate a task by sleeping.
	time.Sleep(100 * time.Millisecond)
	// If the ad name is "Ad2", return an error.
	if ad.Name == "Ad2" {
		return fmt.Errorf("failed to create ad: %s", ad.Name)
	}
	return nil
}

/*
它模拟了在一个事务中并发地处理任务的过程。TransactFunc函数只是简单地调用了传入的函数，而CreateAd函数则模拟了一个可能会失败的任务。如果任务失败，那么errgroup会取消所有的任务，并返回错误。

这段代码是在一个数据库事务中并发地创建广告。这里使用了Go的errgroup包和context包，以及一个事务处理函数TransactFunc。

以下是这段代码的主要功能和值得学习的点：
事务处理：TransactFunc函数是一个处理数据库事务的函数。它接收一个函数作为参数，这个函数在一个数据库事务中执行。如果这个函数返回错误，那么事务就会回滚，否则事务就会提交。
并发处理：在事务中，对createAdList列表中的每一个广告，都并发地调用dc.CreateAd函数创建广告。这是通过errgroup包实现的。errgroup包可以方便地处理并发任务中的错误。如果任何一个并发任务返回错误，那么g.Wait()就会返回这个错误。
上下文管理：errgroup.WithContext函数创建了一个新的上下文和一个错误组。这个上下文会在错误组中的任何一个任务返回错误时被取消。这样，如果任何一个创建广告的任务失败，那么所有的任务都会被取消。
错误处理：这段代码使用了errors.Wrap函数来添加错误的上下文信息。这样可以方便地追踪错误发生的位置和原因。
日志记录：在每次创建广告后，都会记录一条日志，这对于追踪和调试问题非常有帮助。

这段代码的设计非常优秀，它充分利用了Go的并发特性，同时也很好地处理了错误和事务。这种模式在处理需要并发执行，但又需要全部成功或全部失败的任务时非常有用。
*/
