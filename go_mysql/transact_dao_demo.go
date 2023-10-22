package main

import (
	"context"
	"fmt"
)

type Dao interface {
	InsertData(ctx context.Context) error
}

type BaseDao struct{}

func (d *BaseDao) InsertData(ctx context.Context) error {
	fmt.Println("Inserting data...")
	// 模拟插入数据成功
	return nil
}

type dao struct {
	BaseDao
}

type TransactFunc func(ctx context.Context, dao Dao) error

func (f TransactFunc) Transact(ctx context.Context, dao Dao) error {
	return f(ctx, dao)
}

func (d *dao) Transact(ctx context.Context, f TransactFunc) error {
	fmt.Println("Starting transaction...")
	//err := f(ctx, d)
	err := f.Transact(ctx, d)
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
		return dao.InsertData(ctx)
	})
	err := d.Transact(ctx, f)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction succeeded.")
	}
}

/*
我们可以将事务的开始、提交和回滚逻辑封装在Transact方法中，然后在TransactFunc方法中传入具体的业务逻辑。
这样，业务逻辑和事务处理逻辑就被解耦了，使得代码更加清晰，易于维护。
同时，由于TransactFunc是一个函数类型，我们可以根据需要动态地传入不同的业务处理函数，增加了代码的灵活性。
*/
