package main

import (
	"context"
	"errors"
	"fmt"
)

type Dao interface {
	InsertData(ctx context.Context, data string) error
}

type baseDao struct{}

func (d *baseDao) InsertData(ctx context.Context, data string) error {
	fmt.Println("Inserting data:", data)
	// 模拟插入数据失败
	if data == "bad" {
		return errors.New("insert failed")
	}
	return nil
}

type dao struct {
	baseDao
}

type TransactFunc func(ctx context.Context, dao Dao) error

func (d *dao) Transact(ctx context.Context, f TransactFunc) error {
	fmt.Println("Starting transaction...")
	err := f(ctx, d)
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
	f := TransactFunc(func(ctx context.Context, dao Dao) error {
		if err := dao.InsertData(ctx, "good"); err != nil {
			return err
		}
		if err := dao.InsertData(ctx, "bad"); err != nil {
			return err
		}
		return nil
	})
	d.Transact(ctx, f)
}

/*
接口的使用：在代码中Dao 是一个接口，这意味着你可以传入任何实现了 Dao 接口的对象。
这样，你就可以在事务中执行这个对象的 InsertData 方法。这提供了很大的灵活性，因为你可以在一个事务中插入任何你想要的数据，只要这个数据的类型实现了 Dao 接口。

函数类型的使用：你的代码中定义了一个函数类型 TransactFunc，并且 Transact 方法接受一个 TransactFunc 类型的参数。
这样，你可以在一个事务中执行任何满足 TransactFunc 类型的函数。这提供了很大的灵活性，因为你可以在一个事务中执行任何你想要的操作，只要这个操作的函数满足 TransactFunc 类型。


总结:
Dao 接口使得 TransactFunc 可以接受任何实现了 Dao 接口的对象，这提供了很大的灵活性。
TransactFunc 函数类型使得 Transact 方法可以接受任何满足 TransactFunc 类型的函数，这也提供了很大的灵活性。
*/
