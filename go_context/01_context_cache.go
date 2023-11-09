package main

import (
	"context"
	"fmt"
	"sync"
)

type CacheKey string

var catcheKey CacheKey = "context_cache" //使用接口级别换成,小写使其包内可见

type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.store[key]
	return value, ok
}

func WithCache(ctx context.Context, cache *Cache) context.Context {
	return context.WithValue(ctx, catcheKey, cache)
}

func GetCache(ctx context.Context) *Cache {
	cache, _ := ctx.Value(catcheKey).(*Cache)
	return cache
}

func main() {
	cache := NewCache()
	cache.Set("key1", "value1")
	cache.Set("key2", 42)

	ctx := WithCache(context.Background(), cache)

	c := GetCache(ctx)
	value1, _ := c.Get("key1")
	value2, _ := c.Get("key2")

	fmt.Println("key1:", value1)
	fmt.Println("key2:", value2)
}

/*
方案实现:
定义了一个 Cache 结构体，它包含一个读写锁和一个存储键值对的 map。提供了 Set 和 Get 方法来设置和获取缓存中的值。
为了将缓存与上下文关联，我们定义了 WithCache 和 GetCache 函数。WithCache 函数将缓存添加到上下文中，而 GetCache 函数从上下文中获取缓存。

这种基于上下文的缓存设计的优点主要有以下几点：

1. 线程安全：通过使用读写锁，这种设计可以在多线程环境下安全地读写缓存。

2. 灵活性：由于缓存值的类型为 interface{}，这意味着你可以将任何类型的值存储在缓存中，这提供了很大的灵活性。

3. 上下文关联：通过将缓存与上下文关联，你可以在处理请求的过程中方便地访问缓存。这对于在处理请求的过程中需要共享数据的场景非常有用。

这种设计主要用在需要在处理请求的过程中共享数据的场景中。例如，你可能需要在处理用户请求的过程中缓存一些数据，以避免重复的数据库查询。
通过将缓存与上下文关联，你可以在处理请求的过程中方便地访问这些数据。
*/
