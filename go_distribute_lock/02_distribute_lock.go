package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	// redisDistributedLockKey redis key
	redisDistributedLockKey = "my-key"
	// keyExpiredTime 锁的过期时间
	keyExpiredTime = time.Second * 20
)

func main() {
	ctx := context.Background()
	// Connect to redis.
	client := redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	})
	defer client.Close()
	// 获取锁
	isSet, err := client.SetNX(ctx, redisDistributedLockKey, uuid.New().String(), keyExpiredTime).Result()
	if err != nil {
		log.Fatalf("setnx key[%v] failed, err:%v", redisDistributedLockKey, err)
		return
	} else if !isSet {
		log.Fatalf("obtain redis key[%v] failed", redisDistributedLockKey)
		return
	}
	fmt.Printf("obtain redis lock succ, key:%v\n", redisDistributedLockKey)
	// 执行业务逻辑 do something
	doBusiLogic()
	time.Sleep(10 * time.Second)
	// 释放锁
	client.Del(ctx, redisDistributedLockKey)
	fmt.Printf("release lock succ, key:%v\n", redisDistributedLockKey)
}

func doBusiLogic() {
	fmt.Println("do business logic")
}

// 问题：1. 死锁  机器a获取到锁之后故障导致无法释放,解决办法：锁设置过期时间
// 2. 锁错误释放 机器a获取锁之后在处理业务逻辑时阻塞,锁过期释放，机器b获取锁之后被复活的机器a释放锁
// 解决办法: 本质原因是在client在释放锁时没有去判断当前锁是不是属于自己。
// 所以这里可以通过给每个锁加一个token，每个client释放锁时判断当前锁是否是自己的锁，如果是才能释放。
