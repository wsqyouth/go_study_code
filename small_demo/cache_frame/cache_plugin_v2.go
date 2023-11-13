package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
)

type Cache interface {
	Set(ctx context.Context, key string, data interface{})
	Get(ctx context.Context, key string, data interface{}) bool
}

type RedisCache struct {
	client string
	data   []byte //暂时存储,真实场景下从数据源获取数据
}

func (r *RedisCache) Set(ctx context.Context, key string, data interface{}) {
	fmt.Println("Setting value in Redis cache")
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		fmt.Println("Error encoding data:", err)
		return
	}
	r.data = buf.Bytes()
}

func (r *RedisCache) Get(ctx context.Context, key string, data interface{}) bool {
	fmt.Println("Getting value from Redis cache")
	if r.data == nil {
		return false
	}
	dec := gob.NewDecoder(bytes.NewReader(r.data))
	err := dec.Decode(data)
	if err != nil {
		fmt.Println("Error decoding data:", err)
		return false
	}
	return true
}

func main() {
	ctx := context.Background()

	// redis全部使用指针用法
	redis := RedisCache{} // 这里使用饿汉式直接初始化,主要学习不同的用法
	fmt.Println("Redis string pointer demo:")
	demoStringPointer(ctx, &redis)
	demoStructPointer(ctx, &redis)

}

type MyStruct struct {
	Field1 string
	Field2 int
}

func demoStringPointer(ctx context.Context, cache Cache) {

	str := "value"
	cache.Set(ctx, "key", str)
	var val string
	if cache.Get(ctx, "key", &val) {
		fmt.Println(val)
	}
}

func demoStructPointer(ctx context.Context, cache Cache) {

	myStruct := MyStruct{"Hello", 123}
	cache.Set(ctx, "key", myStruct)
	var val MyStruct
	if cache.Get(ctx, "key", &val) {
		fmt.Println(val)
	}
}
