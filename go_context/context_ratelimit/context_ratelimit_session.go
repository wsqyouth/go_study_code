package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/time/rate"
)

func main() {
	// 创建一个新的DBPool
	dbPool, err := NewDBPool("user:password@/dbname")
	if err != nil {
		log.Fatalf("Failed to create DBPool: %v", err)
	}
	// 创建一个新的SessionPool
	sessionPool := NewSessionPool(dbPool)
	// 非事务操作
	session.Select("SELECT * FROM users", handler)
	// 开始一个新的事务
	session.Begin()
	// 事务操作
	session.Insert("INSERT INTO users (name) VALUES (?)", "Alice")
	session.Update("UPDATE users SET age = ? WHERE name = ?", 25, "Alice")
	// 提交事务
	session.Commit()
}

type key int

const (
	contextRateLimitKey key = 1
)

// RateLimitKey 限频key
type RateLimitKey struct {
	DataSource string
}

type DBPool struct {
	pool         map[string]*sql.DB
	rateLimiters map[string]*rate.Limiter
	mu           sync.RWMutex
}

func NewDBPool() *DBPool {
	return &DBPool{
		pool:         make(map[string]*sql.DB),
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

// DBPoolConfig 数据库连接配置
type DBPoolConfig struct {
	key    string
	source string
}

// 根据请求确定key和对应的账号密码
func getDBConfig(ctx context.Context, query string) *DBPoolConfig {
	fmt.Println(query)
	return &DBPoolConfig{
		key:    "127.0.0.1:3306:sql_test",
		source: "coopers:2019Youth@tcp(127.0.0.1:3306)/sql_test?charset=utf8mb4&parseTime=True",
	}
}

func (p *DBPool) getDb(ctx context.Context, source string) (*sql.DB, error) {
	var err error
	p.mu.RLock()
	db, isExist := p.pool[source]
	p.mu.RUnlock()
	if !isExist {
		fmt.Println(source)
		db, err = sql.Open("mysql", source)
		if err != nil {
			return nil, err
		}
		p.mu.Lock()
		p.pool[source] = db
		p.mu.Unlock()
	}
	return db, nil
}

// getRateLimter 首先尝试获取读锁来检查限制器是否存在。如果不存在，我们释放读锁，获取写锁，然后创建和存储新的限制器。
func (p *DBPool) getRateLimter(ctx context.Context, key string) *rate.Limiter {
	p.mu.RLock()
	limiter, isExist := p.rateLimiters[key]
	p.mu.RUnlock()
	if !isExist {
		p.mu.Lock()
		limiter = rate.NewLimiter(rate.Every(time.Second), 1) // qps=1
		p.rateLimiters[key] = limiter
		p.mu.Unlock()
	}
	return limiter
}

type RowHandler func(*sql.Rows) error

// executeQuery 执行sql, 存在限频器则进行限频
func (p *DBPool) executeQuery(ctx context.Context, dbConfig *DBPoolConfig, query string, handler RowHandler) error {
	limitKey, ok := getContextRateLimit(ctx)
	if ok {
		limiter := p.getRateLimter(ctx, limitKey.DataSource)
		if !limiter.Allow() {
			return fmt.Errorf("query limit")
		}
	}
	// Execute the query...
	db, err := p.getDb(ctx, dbConfig.source)
	if err != nil {
		return fmt.Errorf("get database err: %v", err)
	}
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	if handler != nil {
		if err := handler(rows); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}
	fmt.Println("executing query: ", query)
	return nil
}

func (p *DBPool) Insert(ctx context.Context, dbConfig *DBPoolConfig, query string, args ...interface{}) error {
	return p.executeQuery(ctx, dbConfig, query, nil)
}

func (p *DBPool) Update(ctx context.Context, dbConfig *DBPoolConfig, query string, args ...interface{}) error {
	return p.executeQuery(ctx, dbConfig, query, nil)
}

func (p *DBPool) Delete(ctx context.Context, dbConfig *DBPoolConfig, query string, args ...interface{}) error {
	return p.executeQuery(ctx, dbConfig, query, nil)
}

func (p *DBPool) Select(ctx context.Context, dbConfig *DBPoolConfig, query string) error {
	handler := func(rows *sql.Rows) error {
		// process each row
		for rows.Next() {
			var name string
			var age uint
			if err := rows.Scan(&name, &age); err != nil {
				return err
			}
			fmt.Println(name, age)
		}
		return nil
	}
	return p.executeQuery(ctx, dbConfig, query, handler)
}

func setContexRateLimit(ctx context.Context, limitKey RateLimitKey) context.Context {
	return context.WithValue(ctx, contextRateLimitKey, limitKey)
}

func getContextRateLimit(ctx context.Context) (limitKey RateLimitKey, ok bool) {
	limitKey, ok = ctx.Value(contextRateLimitKey).(RateLimitKey)
	return
}

type Session struct {
	db *sql.DB
	tx *sql.Tx
}

func (s *Session) Begin() error {
	var err error
	s.tx, err = s.db.Begin()
	return err
}

func (s *Session) Commit() error {
	if s.tx == nil {
		return errors.New("no transaction started")
	}
	err := s.tx.Commit()
	s.tx = nil // reset transaction after commit
	return err
}

func (s *Session) Rollback() error {
	if s.tx == nil {
		return errors.New("no transaction started")
	}
	err := s.tx.Rollback()
	s.tx = nil // reset transaction after rollback
	return err
}

func (s *Session) executeQuery(query string, handler RowHandler) error {
	var rows *sql.Rows
	var err error

	if s.tx != nil {
		rows, err = s.tx.Query(query)
	} else {
		rows, err = s.db.Query(query)
	}

	if err != nil {
		return err
	}
	defer rows.Close()

	if handler != nil {
		if err := handler(rows); err != nil {
			return err
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}

func (s *Session) Insert(query string, args ...interface{}) error {
	return s.executeQuery(query, nil)
}

func (s *Session) Update(query string, args ...interface{}) error {
	return s.executeQuery(query, nil)
}

func (s *Session) Delete(query string, args ...interface{}) error {
	return s.executeQuery(query, nil)
}

func (s *Session) Select(query string, handler RowHandler) error {
	return s.executeQuery(query, handler)
}

type SessionPool struct {
	dbPool   *DBPool
	sessions map[string]*Session
}

func NewSessionPool(dbPool *DBPool) *SessionPool {
	return &SessionPool{
		dbPool:   dbPool,
		sessions: make(map[string]*Session),
	}
}

func (p *SessionPool) GetSession(sessionID string) (*Session, error) {
	session, ok := p.sessions[sessionID]
	if ok {
		return session, nil
	}

	db := p.dbPool.pool
	session = &Session{db: db}
	p.sessions[sessionID] = session
	return session, nil
}

func (p *SessionPool) ReleaseSession(sessionID string) {
	delete(p.sessions, sessionID)
}

// DBPool管理数据库连接，SessionPool管理Session，并通过DBPool获取数据库连接。每个Session都有一个唯一的sessionID，你可以通过这个sessionID从SessionPool中获取或释放Session。
