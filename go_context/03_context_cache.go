package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

type CacheKey string

type cacheItem struct {
	ret  interface{}
	err  error
	once sync.Once
}

type Cache struct {
	m    map[string]*cacheItem
	lock sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]*cacheItem),
	}
}

// func (c *Cache) Set(ctx context.Context, key string, value interface{}, err error) {
// 	c.lock.Lock()
// 	defer c.lock.Unlock()
// 	c.m[key] = &cacheItem{ret: value, err: err}
// }

// sync.Once 的确没有被使用。sync.Once 通常用于确保某个操作只执行一次，这在初始化或者延迟加载等场景中非常有用。
func (c *Cache) Set(ctx context.Context, key string, value interface{}, err error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	item, ok := c.m[key]
	if !ok {
		item = &cacheItem{}
		c.m[key] = item
	}
	item.once.Do(func() {
		item.ret = value
		item.err = err
	})
}

func (c *Cache) Get(ctx context.Context, key string) (interface{}, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, ok := c.m[key]
	if !ok {
		return nil, errors.New("key not found")
	}
	return item.ret, item.err
}

func WithCache(ctx context.Context, cache *Cache) context.Context {
	return context.WithValue(ctx, CacheKey("cache"), cache)
}

func GetCache(ctx context.Context) *Cache {
	cache, _ := ctx.Value(CacheKey("cache")).(*Cache)
	return cache
}

func Get(ctx context.Context, key string) (interface{}, error) {
	cache := GetCache(ctx)
	if cache == nil {
		return nil, errors.New("cache not found")
	}
	return cache.Get(ctx, key)
}

func Set(ctx context.Context, key string, value interface{}, err error) {
	cache := GetCache(ctx)
	if cache != nil {
		cache.Set(ctx, key, value, err)
	}
}

func main() {
	cache := NewCache()

	ctx := WithCache(context.Background(), cache)

	Set(ctx, "key1", "value1", nil)
	Set(ctx, "key2", 42, nil)

	value1, err1 := Get(ctx, "key1")
	value2, err2 := Get(ctx, "key2")

	fmt.Println("key1:", value1, "error:", err1)
	fmt.Println("key2:", value2, "error:", err2)
}

/*
针对错误无法感知的缺点，我们引入了一个 cacheItem 结构体，它包含了返回值、错误和一个 sync.Once 实例。
Cache 结构体包含了一个存储 cacheItem 的映射。这种设计允许我们在缓存中存储错误信息，从而使得错误更容易被追踪。


这种设计确实解决了错误无法感知的问题，但它可能会使代码变得更复杂。在实际应用中，你需要根据你的需求和场景来权衡这种设计的优缺点。
另外，业界还有其他一些解决方案，例如使用回调函数或者将错误信息作为缓存值的一部分。这些解决方案的优缺点也需要根据具体场景来评估。
*/
