package db

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/pkg/errors"
)

var defaultDBPool DBPool = nil
var once sync.Once

// DBPoolConfig 连接池配置
type DBPoolConfig struct {
	Source  string
	MaxOpen int32
	MaxIdle int32
}

// DBPool 数据库对象池
type DBPool interface {
	Open(ctx context.Context, config *DBPoolConfig) (*sql.DB, error)
	CloseAll()
	GetDBPoolInfo() map[string]*sql.DB
}

type mysqlDBPool struct {
	m     map[string]*sql.DB
	mutex sync.Mutex
}

// GetMySQLDBPool 获取全局唯一的DBPool
func GetMySQLDBPool() DBPool {
	once.Do(func() {
		defaultDBPool = newMysqlDBPool()
	})
	return defaultDBPool
}

func newMysqlDBPool() *mysqlDBPool {
	return &mysqlDBPool{
		m: make(map[string]*sql.DB),
	}
}

// Open 获取连接
func (p *mysqlDBPool) Open(ctx context.Context, config *DBPoolConfig) (*sql.DB, error) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	db, ok := p.m[config.Source]
	if ok {
		return db, nil
	}
	db, err := sql.Open("mysql", config.Source)
	if err != nil {
		return nil, errors.Wrap(err, "open error")
	}

	db.SetMaxOpenConns(int(config.MaxOpen))
	db.SetMaxIdleConns(int(config.MaxIdle))

	// 限制一个连接最大存活时间，加点随机量，避免同时重建连接
	rand.Seed(time.Now().Unix())
	minConnLife := int64(10 * time.Minute)
	maxConnLife := minConnLife + rand.Int63n(minConnLife)
	db.SetConnMaxLifetime(time.Duration(maxConnLife))
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping error")
	}

	p.m[config.Source] = db
	return db, nil
}

// CloseAll 关闭连接
func (p *mysqlDBPool) CloseAll() {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	for _, v := range p.m {
		err := v.Close()
		if err != nil {
			fmt.Printf("close db error, %s", err)
		}
	}
}

// GetDBPoolInfo获取DBPool信息
func (p *mysqlDBPool) GetDBPoolInfo() map[string]*sql.DB {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	return p.m
}
