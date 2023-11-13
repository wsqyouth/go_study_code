package main

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"time"
)

type Cache interface {
	Set(ctx context.Context, key string, data interface{}, d time.Duration)
	Get(ctx context.Context, key string, data interface{}) bool
}

type RedisCache struct {
	client string
	data   interface{} //暂时存储,真实场景下从数据源获取数据
}

func (r *RedisCache) Set(ctx context.Context, key string, data interface{}, d time.Duration) {
	fmt.Println("Setting value in Redis cache")
	r.data = data
}

func (r *RedisCache) Get(ctx context.Context, key string, data interface{}) bool {
	fmt.Println("Getting value from Redis cache")

	if r.data == nil {
		return false
	}
	val := reflect.ValueOf(data)
	if val.Kind() == reflect.Ptr {
		val.Elem().Set(reflect.ValueOf(r.data).Elem())
		return true
	} else if val.Kind() == reflect.String {
		// Assuming r.data is of type string
		strData, ok := r.data.(string)
		if !ok {
			return false
		}
		val.SetString(strData)
		return true
	}

	return false
}

type MemcacheCache struct {
	client string
}

func (m *MemcacheCache) Set(ctx context.Context, key string, data interface{}, d time.Duration) {
	fmt.Println("Setting value in Memcache cache")
}

func (m *MemcacheCache) Get(ctx context.Context, key string, data interface{}) bool {
	fmt.Println("Getting value from Memcache cache")
	return true
}

var (
	redisCache    *RedisCache
	memcacheCache *MemcacheCache
	onceRedis     sync.Once
	onceMemcache  sync.Once
)

func GetRedisInstance() *RedisCache {
	onceRedis.Do(func() {
		redisCache = &RedisCache{}
	})
	return redisCache
}

func GetMemcacheInstance() *MemcacheCache {
	onceMemcache.Do(func() {
		memcacheCache = &MemcacheCache{}
	})
	return memcacheCache
}

func main() {
	ctx := context.Background()

	// memcache演示基本的用法
	memcache := GetMemcacheInstance()
	memcache.Set(ctx, "key", "value", 1*time.Second)
	memcache.Get(ctx, "key", "value")

	// redis全部使用指针用法
	redis := GetRedisInstance()
	fmt.Println("Redis string pointer demo:")
	demoStringPointer(ctx, redis)
	demoStructPointer(ctx, redis)

}

type MyStruct struct {
	Field1 string
	Field2 int
}

func demoStringPointer(ctx context.Context, cache Cache) {
	str := "value"
	cache.Set(ctx, "key", &str, 0)
	var val string
	if cache.Get(ctx, "key", &val) {
		fmt.Println(val)
	}
}

func demoStructPointer(ctx context.Context, cache Cache) {

	myStruct := MyStruct{"Hello", 123}
	cache.Set(ctx, "key", &myStruct, 0)
	var val MyStruct
	if cache.Get(ctx, "key", &val) {
		fmt.Println(val)
	}
}
