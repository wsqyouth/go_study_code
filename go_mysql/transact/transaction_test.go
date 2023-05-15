// dao_test.go
package main

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

func TestTransactFunc(t *testing.T) {
	baseDao := &dao{}
	d := NewDao(baseDao)

	f := func(ctx context.Context, dao Dao) error {
		return nil
	}

	err := d.TransactFunc(context.Background(), f)
	assert.NoError(t, err)
}

func TestTransactFuncWithError(t *testing.T) {
	baseDao := &dao{}
	d := NewDao(baseDao)

	f := func(ctx context.Context, dao Dao) error {
		return errors.New("transaction error")
	}

	err := d.TransactFunc(context.Background(), f)
	assert.Error(t, err)
}

func TestTransact(t *testing.T) {
	baseDao := &dao{}
	d := NewDao(baseDao)

	transactor := TransactFunc(func(ctx context.Context, dao Dao) error {
		return nil
	})

	err := d.Transact(context.Background(), transactor)
	assert.NoError(t, err)
}

func TestTransactWithError(t *testing.T) {
	baseDao := &dao{}
	d := NewDao(baseDao)

	transactor := TransactFunc(func(ctx context.Context, dao Dao) error {
		return errors.New("transaction error")
	})

	err := d.Transact(context.Background(), transactor)
	assert.Error(t, err)
}

func TestTransactWithGomonkey(t *testing.T) {
	baseDao := &dao{}
	d := NewDao(baseDao)

	transactor := TransactFunc(func(ctx context.Context, dao Dao) error {
		return nil
	})

	patches := gomonkey.ApplyMethod(reflect.TypeOf(baseDao), "TransactFunc", func(_ *dao, _ context.Context, _ TransactFunc) error {
		return nil
	})
	defer patches.Reset()

	err := d.Transact(context.Background(), transactor)
	assert.NoError(t, err)
}
