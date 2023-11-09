package main

import (
	"context"
	"fmt"
	"sync"
)

type CacheKey string

// 接口类型定义
type ContextCache interface {
	Get(ctx context.Context, key string) interface{}
	Set(ctx context.Context, key string, value interface{})
}

type Cache struct {
	mu    sync.RWMutex
	store map[string]interface{}
}

func NewCache() *Cache {
	return &Cache{
		store: make(map[string]interface{}),
	}
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = value
}

func (c *Cache) Get(ctx context.Context, key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value := c.store[key]
	return value
}

// 接口类型传参
func WithCache(ctx context.Context, cache ContextCache) context.Context {
	return context.WithValue(ctx, CacheKey("cache"), cache)
}

// 接口类型断言
func GetCache(ctx context.Context) ContextCache {
	cache, _ := ctx.Value(CacheKey("cache")).(ContextCache)
	return cache
}

// 全局函数Get
func Get(ctx context.Context, key string) interface{} {
	cache := GetCache(ctx)
	if cache == nil {
		return nil
	}
	return cache.Get(ctx, key)
}

// 全局函数Set
func Set(ctx context.Context, key string, value interface{}) {
	cache := GetCache(ctx)
	if cache != nil {
		cache.Set(ctx, key, value)
	}
}

func main() {
	cache := NewCache()

	ctx := WithCache(context.Background(), cache)

	Set(ctx, "key1", "value1")
	Set(ctx, "key2", 42)

	value1 := Get(ctx, "key1")
	value2 := Get(ctx, "key2")

	fmt.Println("key1:", value1)
	fmt.Println("key2:", value2)
}

/*
使用接口对外封装:
Get 函数是一个全局函数，它从上下文中获取缓存并返回指定键的值。
这种设计的主要优点是简化了从上下文中获取缓存和值的操作，你只需要调用一个函数就可以完成这些操作，而不需要手动从上下文中获取缓存并调用其 Get 方法。
这种设计的缺点是它可能会隐藏一些错误。
例如，如果上下文中没有缓存，Get 函数会返回 nil，而不是返回一个错误。
这可能会使得错误更难以追踪，因为你可能不知道 Get 函数返回 nil 的原因是因为上下文中没有缓存，还是因为缓存中没有指定的键。
*/
