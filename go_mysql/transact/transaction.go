// dao.go
package main

import (
	"context"
)

type Dao interface {
	TransactFunc(ctx context.Context, f TransactFunc) error
}

type BaseDao struct {
	DbName string
}

type dao struct {
	BaseDao Dao
}

func NewDao(baseDao Dao) Dao {
	return &dao{BaseDao: baseDao}
}
func (dao Dao) Transact(ctx context.Context, transactor Transactor) (err error) {
	return transactor.Transact(ctx)
}

type Transactor interface {
	Transact(ctx context.Context, dao Dao) error
}

type TransactFunc func(ctx context.Context, dao Dao) (err error)

func (f TransactFunc) Transact(ctx context.Context, dao Dao) error {
	return f(ctx, dao)
}

func (d *dao) TransactFunc(ctx context.Context, f TransactFunc) error {
	return d.Transact(ctx, f)
}

func (d *dao) Transact(ctx context.Context, transactor Transactor) error {
	return d.BaseDao.TransactFunc(ctx, func(ctx context.Context, dao Dao) (err error) {
		return transactor.Transact(ctx, d)
	})
}
