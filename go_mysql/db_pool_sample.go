package main

import (
	"database/sql"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type dbPoolConfig struct {
	Key     string
	Source  string
	MaxOpen int32
	MaxIdle int32
}

type dbPool struct {
	config dbPoolConfig
	pool   *sync.Pool
	conns  []*sql.DB
	mu     sync.Mutex
	closed bool
}

func newDBPool(config dbPoolConfig) (*dbPool, error) {
	pool := &sync.Pool{
		New: func() interface{} {
			db, err := sql.Open("mysql", config.Source)
			if err != nil {
				return nil
			}
			db.SetMaxOpenConns(int(config.MaxOpen))
			db.SetMaxIdleConns(int(config.MaxIdle))
			db.SetConnMaxLifetime(time.Minute * 5)
			return db
		},
	}

	// Test the connection
	db := pool.Get().(*sql.DB)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	pool.Put(db)
	return &dbPool{
		config: config,
		pool:   pool,
		conns:  []*sql.DB{db},
	}, nil
}

func (p *dbPool) Get() (*sql.DB, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return nil, errors.New("pool is closed")
	}
	db := p.pool.Get().(*sql.DB)
	if db == nil {
		return nil, errors.New("failed to get a connection from the pool")
	}
	p.conns = append(p.conns, db)
	return db, nil
}

func (p *dbPool) Put(db *sql.DB) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if !p.closed {
		p.pool.Put(db)
	}
}

func (p *dbPool) Close() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.closed {
		return errors.New("pool is already closed")
	}

	var firstErr error
	for _, db := range p.conns {
		if err := db.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	p.closed = true
	return firstErr
}

func main() {
	config := dbPoolConfig{
		Key:     "test",
		Source:  "coopers:2019Youth@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True",
		MaxOpen: 10,
		MaxIdle: 5,
	}

	pool, err := newDBPool(config)
	if err != nil {
		fmt.Println("Error creating pool:", err)
		return
	}
	defer pool.Close()

	db, err := pool.Get()
	if err != nil {
		fmt.Println("Error getting connection:", err)
		return
	}
	// 打印db.State
	fmt.Println(db.Stats())
	// Use the connection for a query
	rows, err := db.Query("SELECT name,age FROM user")
	if err != nil {
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Process the query results
	for rows.Next() {
		var name string
		var age uint
		if err := rows.Scan(&name, &age); err != nil {
			fmt.Println("Error scan:", err)
			return
		}
		fmt.Println(name, age)
	}

	pool.Put(db)
}
