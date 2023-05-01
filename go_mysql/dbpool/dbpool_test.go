package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/agiledragon/gomonkey"
	"github.com/stretchr/testify/assert"
)

func Test_mysqlDBPool_Open(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()
	testKey := "testKey123"
	p := &mysqlDBPool{
		m: map[string]*sql.DB{
			testKey: db,
		},
	}
	config := &DBPoolConfig{
		Source:  testKey,
		MaxOpen: 456,
		MaxIdle: 123,
	}
	gotDB, err := p.Open(context.Background(), config)
	assert.Nil(t, err)
	assert.Equal(t, gotDB, db)

	testNewKey := "testNewKey123"
	config = &DBPoolConfig{
		Source:  testNewKey,
		MaxOpen: 456,
		MaxIdle: 123,
	}
	patches := gomonkey.ApplyFunc(sql.Open, func(driverName, dataSourceName string) (*sql.DB, error) {
		return db, nil
	})
	defer patches.Reset()
	gotDB2, err := p.Open(context.Background(), config)
	assert.Nil(t, err)
	assert.Equal(t, int32(gotDB2.Stats().MaxOpenConnections), config.MaxOpen)
}

func Test_mysqlDBPool_Get(t *testing.T) {
	pool := GetMySQLDBPool()
	pool1 := GetMySQLDBPool()
	pool2 := GetMySQLDBPool()
	assert.Equal(t, pool, pool1)
	assert.Equal(t, pool, pool2)
	pool.CloseAll()
}

func Test_GetDBPoolInfo(t *testing.T) {
	pool := GetMySQLDBPool()
	poolInfo := pool.GetDBPoolInfo()
	assert.NotNil(t, poolInfo)
}
