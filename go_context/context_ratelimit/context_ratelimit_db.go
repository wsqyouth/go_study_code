package main

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()
	query := "SELECT name,age FROM user limit 3"
	dbConfig := getDBConfig(ctx, query)
	ctx = setContexRateLimit(ctx, RateLimitKey{dbConfig.key})
	dbPool := NewDBPool()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := dbPool.Select(ctx, dbConfig, query)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			fmt.Println("Query executed success. num:", i)
		}(i)
	}
	wg.Wait()
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

/*
总结：
功能实现分为三个阶段：
1. 通过conext val实现对某数据源的限频操作注入,完成基于context val的读写操作
2. 添加限频器,当ctx要求对数据源限频时实现qps=1的限频
3. 使用waitGroup进行并发测试，检查限频功能是否生效
4. 并发场景下，对dbpool中map的读写需要考虑锁的场景,添加读写锁
5. 抽象提取dbConfig,利用key实现对某个数据库实例进行限频,利用source存储对应数据库的账号密码
6. 抽象executeQuery,封装出CRUD,方便使用者调用
*/
