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
	keyExpiredTime = time.Second * 11
	retryTimes     = 3
	redisAddr      = "127.0.0.1:6379"
)

var (
	// luaRelease 释放锁的lua脚本
	luaRelease = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)
	// luaRefresh 续期的lua脚本
	luaRefresh = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("pexpire", KEYS[1], ARGV[2]) else return 0 end`)
)

// RedisLock redis 分布式锁
type RedisLock struct {
	Key        string
	Value      string
	ExpireTime time.Duration
	*redis.Client
	// GoroutineNum 当前启动的goroutine编号，方便调试
	GoroutineNum int
	StopBusi     chan struct{}
	Done         chan struct{}
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
		StopBusi:     make(chan struct{}),
		Done:         make(chan struct{}),
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
	lock.Done <- struct{}{}
	res, err := luaRelease.Run(context.Background(), lock.Client, []string{lock.Key}, lock.Value).Result()
	if err == redis.Nil {
		log.Printf("goroutine%d, release redis lock failed, key[%v], value[%v], redis.Nil, res:%v",
			lock.GoroutineNum, lock.Key, lock.Value, res)
		return fmt.Errorf("goroutine%d release redis lock failed", lock.GoroutineNum)
	} else if err != nil {
		log.Printf("goroutine%d, release redis lock expected, key[%v], value[%v], err:%v",
			lock.GoroutineNum, lock.Key, lock.Value, err)
		return err
	}

	if i, ok := res.(int64); !ok || i != 1 {
		log.Printf("goroutine%d, release redis lock failed, key[%v], value[%v], retcode invalide:%v",
			lock.GoroutineNum, lock.Key, lock.Value, i)
		return fmt.Errorf("goroutine%d release redis lock failed", lock.GoroutineNum)
	}

	log.Printf("goroutine%d release redis lock succ, key:%s, value:%s\n", lock.GoroutineNum, lock.Key, lock.Value)
	return nil
}

/*Refresh 续期
业务逻辑执行完 停止锁续期
业务逻辑未执行完
	续期成功，则继续续期
	续期异常，中止当前业务逻辑执行
	续期过程锁被释放，中止当前业务逻辑执行
*/
func (lock *RedisLock) Refresh(ctx context.Context) error {
	log.Printf("refresh")
	// 提前1s续期
	interval := lock.ExpireTime - time.Second
	var pexpireCount int
	fmt.Println(interval.Seconds())
	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	for {
		log.Println("refresh start")
		select {
		case <-lock.Done:
			return nil
		case <-ticker.C:
			pexpireCount++
			status, err := luaRefresh.Run(ctx, lock.Client, []string{lock.Key}, lock.Value, lock.ExpireTime.Milliseconds()).Result()
			if err != nil { // 续期异常，中止当前业务逻辑执行
				log.Printf("goroutine%d value[%s] refresh failed, err:%v", lock.GoroutineNum, lock.Value, err)
				lock.StopBusi <- struct{}{}
				return fmt.Errorf("goroutine%d value[%s] refresh failed, err:%v", lock.GoroutineNum, lock.Value, err)
			} else if status == int64(1) { // 续期成功
				log.Printf("goroutine%d value[%s] %dth pexpire succ value", lock.GoroutineNum, lock.Value, pexpireCount)
			} else { // 续期过程中锁被释放，中止当前业务逻辑执行
				log.Printf("goroutine%d not acquire lock value[%s], stop busi", lock.GoroutineNum, lock.Value)
				lock.StopBusi <- struct{}{}
				return fmt.Errorf("goroutine%d not acquire lock value[%s], stop busi", lock.GoroutineNum, lock.Value)
			}
		}
	}
}

func doBusiLogic(lock *RedisLock) {
	log.Printf("goroutine%d, do business logic", lock.GoroutineNum)
	for i := 0; i < 20; i++ {
		log.Printf("do busi logic i:%d", i+1)
		time.Sleep(time.Second * 2)
	}
}

// Example ...
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

	// 启动锁进行续期
	go lock.Refresh(ctx)

	defer lock.Release(ctx)
	doBusiLogic(lock)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Example(&wg, i+1)
		time.Sleep(40 * time.Second)
	}
	wg.Wait()
}
