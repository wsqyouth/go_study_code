package main

import (
	"fmt"
	"sync"
)

type Singleton struct{}

var (
	singletonInstance *Singleton
	once              sync.Once
)

// 通过once.Do()确保instance只会被初始化一次
func GetSingletonInstance() *Singleton {
	once.Do(func() {
		singletonInstance = &Singleton{}
	})
	return singletonInstance
}

func main() {
	singletonInstance := GetSingletonInstance()
	fmt.Println(singletonInstance)
	fmt.Println("Hello World")
}

/*
官方：https://pkg.go.dev/sync#Once.Do

单例模式：确保全局只有一个实例对象，避免重复创建资源。
延迟初始化：在程序运行过程中需要用到某个资源时，通过 sync.Once 动态地初始化该资源。
只执行一次的操作：例如只需要执行一次的配置加载、数据清理等操作。

源码实现：
首先它会通过原子操作atomic.LoadUint32（保证并发安全） 检查 done 的值，如果为 0，表示 f 函数没有被执行过，然后执行 doSlow 方法。
在 doSlow 方法里，首先对互斥锁 m 进行加锁，确保在多个协程访问时，只有一个协程能执行 f 函数。接着再次检查 done 变量的值，
如果 done 的值仍为 0，说明 f 函数没有被执行过，此时执行 f 函数，最后通过原子操作 atomic.StoreUint32 将 done 变量的值设置为 1。

为什么会有双重检查（double check）的写法
从源码可知，存在两次对 done 的值的判断。
第一次检查：在获取锁之前，先使用原子加载操作 atomic.LoadUint32 检查 done 变量的值，如果 done 的值为 1，表示操作已执行，此时直接返回，不再执行 doSlow 方法。这一检查可以避免不必要的锁竞争。
第二次检查：获取锁之后，再次检查 done 变量的值，这一检查是为了确保在当前协程获取锁期间，其他协程没有执行过 f 函数。如果 done 的值仍为 0，表示 f 函数没有被执行过。
通过双重检查，可以在大多数情况下避免锁竞争，提高性能。
*/
