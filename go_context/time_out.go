package main

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func main() {

	processTask()
	fmt.Println("-------")
	processTaskNew()

}

func processTask() error {
	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	fmt.Println(time.Now())

	g.Go(func() (e error) {
		err := doWork(ctx)
		if err != nil {
			return err
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}

func processTaskNew() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer cancel()

	ch := make(chan error)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("panic: %#v", err)
				ch <- errors.New("doWork panic")
			}
		}()
		err = doWork(ctx)
		if err != nil {
			fmt.Printf("doWork error: %+v", err)
		}
		ch <- err
	}()

	select {
	case err := <-ch:
		return err
	case <-ctx.Done():
		fmt.Printf("doWork time out")
		return ctx.Err()
	}
}

func doWork(ctx context.Context) error {
	time.Sleep(10 * time.Second)
	fmt.Println("work done")
	return nil
}
