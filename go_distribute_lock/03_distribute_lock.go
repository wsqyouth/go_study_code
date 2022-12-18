package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const (
	// redisDistributedLockKey redis key
	redisDistributedLockKey = "my-key"
	// keyExpiredTime 锁的过期时间
	keyExpiredTime = time.Second * 30
	retryTimes     = 3
	redisAddr      = "127.0.0.1:6379"
)

// RedisLock redis 分布式锁
type RedisLock struct {
	Key        string
	Value      string
	ExpireTime time.Duration
	*redis.Client
	// GoroutineNum 当前启动的goroutine编号，方便调试
	GoroutineNum int
}

// NewRedisLock 生成一个redis分布式锁实例
func NewRedisLock(key string, goroutineNum int) *RedisLock {
	return &RedisLock{
		Key:        redisDistributedLockKey,
		Value:      uuid.New().String(),
		ExpireTime: keyExpiredTime,
		Client: redis.NewClient(&redis.Options{
			Network: "tcp",
			Addr:    redisAddr,
		}),
		GoroutineNum: goroutineNum,
	}
}

// Acquire 获取分布式锁
func (lock *RedisLock) Acquire(ctx context.Context, retryTimes int) error {
	for i := 0; i < retryTimes; i++ {
		isSet, err := lock.Client.SetNX(ctx, lock.Key, lock.Value, lock.ExpireTime).Result()
		if err != nil {
			return fmt.Errorf("lock setnx failed, err:%v", err)
		}

		// 获取锁成功
		if isSet {
			log.Printf("goroutine%d %dth acquire redis lock succ, key:%s, value:%s\n",
				lock.GoroutineNum, i+1, lock.Key, lock.Value)
			return nil
		}
		log.Printf("goroutine%d, acquire %dth retry, key:%s, value:%s\n",
			lock.GoroutineNum, i+1, lock.Key, lock.Value)
		time.Sleep(lock.ExpireTime / 2)
		// time.Sleep(time.Second * 1000000)
	}

	return fmt.Errorf("goroutine%d, acquire redis lock try %v times failed, value:%s",
		lock.GoroutineNum, retryTimes, lock.Value)
}

// Release 释放锁
func (lock *RedisLock) Release(ctx context.Context) error {
	value, _ := lock.Get(ctx, lock.Key).Result()
	// 判断当前锁是不是自己拥有
	if value != lock.Value {
		return fmt.Errorf("release redis lock failed, lock not match, value:%s, lock.Value:%s", value, lock.Value)
	}
	// 释放锁
	lock.Del(ctx, lock.Key)
	log.Printf("goroutine%d release redis lock succ, key:%s, value:%s\n", lock.GoroutineNum, lock.Key, lock.Value)
	return nil
}

func doBusiLogic(lock *RedisLock) {
	log.Printf("goroutine%d, do business logic", lock.GoroutineNum)
}

func Example(wg *sync.WaitGroup, goroutineNum int) {
	log.Printf("start goroutine%v\n", goroutineNum)
	defer wg.Done()
	ctx := context.Background()
	lock := NewRedisLock(redisDistributedLockKey, goroutineNum)
	err := lock.Acquire(ctx, retryTimes)
	if err != nil {
		log.Printf("acquire lock failed, err:%v", err)
		return
	}
	defer lock.Release(ctx)
	doBusiLogic(lock)
	time.Sleep(20 * time.Second)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Example(&wg, i+1)
		time.Sleep(10 * time.Second)
	}
	wg.Wait()
}

// 仍然有问题: 机器a在释放锁是阻塞过期,机器b在获取锁执行时被复活的机器a是否锁，根本原因：获取和释放不是原子的
// 解决方法：保证释放锁的整个操作（get和del）必须是原子的。这里我们通过Lua脚本来保证一致性
