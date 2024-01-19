package main

import (
	"context"
	"errors"
	"fmt"
	"runtime/debug"
)

type HandleFunc func(ctx context.Context, param string) (err error)

func errBeforeMetric(ctx context.Context, h HandleFunc) (err error) {
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("errBeforeMetric panic: %v, error stack trace: %s\n", p, string(debug.Stack()))
			err = errors.New("errBeforeMetric internal error")
		} else if err != nil {
			fmt.Printf("errBeforeMetric error stack trace: %+v\n", err)
		}
	}()

	return h(ctx, "test_param")
}

func main() {
	// 模拟成功的场景
	successHandler := func(ctx context.Context, param string) error {
		fmt.Println("successHandler executed")
		return nil
	}

	// 模拟失败的场景
	errorHandler := func(ctx context.Context, param string) error {
		fmt.Println("errorHandler executed")
		return errors.New("errorHandler error")
	}

	ctx := context.Background()

	// 测试成功的场景
	err := errBeforeMetric(ctx, successHandler)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}

	// 测试失败的场景
	err = errBeforeMetric(ctx, errorHandler)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Success")
	}
}
