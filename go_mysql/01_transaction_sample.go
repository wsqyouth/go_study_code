package main

import (
	"context"
	"errors"
	"fmt"
)

type dao struct{}

func (d *dao) InsertData(ctx context.Context, data string) error {
	fmt.Println("Inserting data:", data)
	// 模拟插入数据失败
	if data == "bad" {
		return errors.New("insert failed")
	}
	return nil
}

/*
简单版本: 接受一个上下文和一个函数作为参数，然后在一个事务中执行这个函数。
这个函数可以是任何你想要的操作，只要它满足 func(ctx context.Context) error 类型。
*/
func (d *dao) Transact(ctx context.Context, f func(ctx context.Context) error) error {
	fmt.Println("Starting transaction...")
	err := f(ctx)
	if err != nil {
		fmt.Println("Rolling back transaction due to error:", err)
		return err
	}
	fmt.Println("Committing transaction...")
	return nil
}

func main() {
	d := &dao{}
	ctx := context.Background()
	f := func(ctx context.Context) error {
		if err := d.InsertData(ctx, "good"); err != nil {
			return err
		}
		if err := d.InsertData(ctx, "bad"); err != nil {
			return err
		}
		return nil
	}
	d.Transact(ctx, f)
}
