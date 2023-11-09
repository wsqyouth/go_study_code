package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type keyType struct{}

var key = keyType{}

func main() {
	ctx := context.WithValue(context.Background(), key, &sync.Map{})

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			Process(ctx, fmt.Sprintf("key%d", i%3)) // 模拟并发访问相同的key
		}(i)
	}
	wg.Wait()
}

func Process(ctx context.Context, cacheKey string) {
	lockMap, ok := ctx.Value(key).(*sync.Map)
	if ok && lockMap != nil {
		tmpLock, _ := lockMap.LoadOrStore(cacheKey, &sync.Mutex{})
		lock, ok := tmpLock.(*sync.Mutex)
		if ok {
			lock.Lock()
			defer lock.Unlock()

			// 模拟处理数据
			fmt.Printf("Start processing %s\n", cacheKey)
			time.Sleep(1 * time.Second)
			fmt.Printf("Finish processing %s\n", cacheKey)
		}
	}
}

/*
这个示例将模拟一个缓存系统，其中使用了 sync.Map 和 sync.Mutex 来确保对每个键的访问是线程安全的。
你会看到，尽管有多个 goroutine 同时访问相同的键，但是每个键只会被处理一次，这就是 sync.Map 和 sync.Mutex 的作用。

sync.Once 只能用于全局的单次初始化，而不能用于每个 key 的单次初始化。这就是为什么这段代码使用了 sync.Map 和 sync.Mutex 来实现每个 key 的单次初始化。


主要目的是实现并发控制，确保在多个 goroutine 同时访问相同的 cache key 时，
只有一个 goroutine 能够进行处理，其他的 goroutine 会等待，直到第一个 goroutine 完成处理并释放锁。
这种设计模式通常被称为 "单次初始化" 或 "只做一次"。

TODO: 还是没太懂,后面再看下 2023年11月09日19:40
*/
