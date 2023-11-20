package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func TryLock(ctx context.Context, resource string, owner string) (func(), error) {
	key := fmt.Sprintf("lock_%s", resource)
	timeout := 5 * time.Second
	if deadline, ok := ctx.Deadline(); ok {
		timeout = time.Until(deadline)
		if timeout <= 0 {
			// 已超时则不用加锁
			return nil, nil
		}
	}
	// 打印检查timeout是否在变化
	fmt.Printf("timeout: %v\n", timeout)
	ok, err := redisClient.SetNX(ctx, key, owner, time.Duration(math.Ceil(timeout.Seconds()))*time.Second).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "lock failed, key: %s", key)
	}
	if !ok {
		var lockBy string
		if err := redisClient.Get(ctx, key).Scan(&lockBy); err == nil && lockBy == owner {
			return nil, nil
		}
		// 打印检查锁被谁持有了
		fmt.Printf("newowner: %v, lockBy:%v\n", owner, lockBy)
		return nil, errors.New(fmt.Sprintf("lock found, key: %s", key))
	}
	return func() {
		if _, err := redisClient.Del(ctx, key).Result(); err != nil {
			fmt.Printf("unlock failed, key: %s, err: %v\n", key, err)
		}
	}, nil
}

var counter int

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			unlock, err := TryLock(ctx, "resource1", fmt.Sprintf("owner%d", id))
			if err != nil {
				fmt.Printf("lock failed: %v\n", err)
				return
			}
			if unlock != nil {
				defer unlock()
			}

			// do something with the locked resource
			counter++
			fmt.Printf("Goroutine %d incremented the counter. Counter: %d\n", id, counter)
		}(i)
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter)
}
