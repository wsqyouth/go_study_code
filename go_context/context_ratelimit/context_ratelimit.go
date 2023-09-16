package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	ctx := context.Background()
	ctx = setContexRateLimit(ctx, RateLimitKey{DataSource: "mysql"})
	dbPool := NewDBPool()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			err := dbPool.executeQuery(ctx, "SELECT * FROM users")
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
	rateLimiters map[string]*rate.Limiter
	mu           sync.RWMutex
}

func NewDBPool() *DBPool {
	return &DBPool{
		rateLimiters: make(map[string]*rate.Limiter),
	}
}

// getRateLimter 首先尝试获取读锁来检查限制器是否存在。如果不存在，我们释放读锁，获取写锁，然后创建和存储新的限制器。
/*
请注意，这种方法可能会导致多个goroutines同时创建限制器，但只有一个会被存储。这通常不是问题，因为rate.NewLimiter的开销很小，但如果你想避免这种情况，你可以使用双重检查锁定模式。
*/
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

// executeQuery 执行sql, 存在限频器则进行限频
func (p *DBPool) executeQuery(ctx context.Context, query string) error {
	limitKey, ok := getContextRateLimit(ctx)
	if ok {
		limiter := p.getRateLimter(ctx, limitKey.DataSource)
		if !limiter.Allow() {
			return fmt.Errorf("query limit")
		}
	}
	// Execute the query...
	// db.Query(query)
	fmt.Println("executing query:------", query)
	return nil
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
*/
