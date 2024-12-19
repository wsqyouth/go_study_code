package main

import (
	"fmt"
	"sync"
)

type Data struct {
	Value string
}

func fetchDataFromRedis(key string) *Data {
	// 模拟从Redis获取数据
	return &Data{Value: "Data for " + key}
}

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	keyList := []string{"key1", "key2", "key3", "key4", "key5"}

	dataMap := make(map[string]*Data, len(keyList))
	sem := make(chan struct{}, 400)

	for _, k := range keyList {
		wg.Add(1)
		sem <- struct{}{} // 获取一个槽位
		go func(k string) {
			defer wg.Done()
			defer func() { <-sem }()      // 释放槽位
			data := fetchDataFromRedis(k) // 执行Redis查询逻辑，并转换为Data结构指针
			mu.Lock()
			dataMap[k] = data
			mu.Unlock()
		}(k)
	}

	wg.Wait()

	ret := make([]*Data, 0)
	for _, k := range keyList {
		ret = append(ret, dataMap[k])
	}

	// 打印结果
	for _, data := range ret {
		fmt.Println(data.Value)
	}
}

/*
旧方案：
 先通过 keyList 拉取一个长度超过 4万的列表，然后从 Redis 中串行拉取列表中每项的信息，耗时 4 min
新方案：
 使用 goroutine 并发的去请求 Redis。但是因为怕并发过高带来副作用，所以需要限制并发数量，就使用带缓冲的 channel 控制
优点：
并发性能提升 ：通过使用goroutine并发地从Redis查询数据，显著提高了数据获取的速度，特别是当涉及到大量key时。
资源控制 ：通过使用信号量sem限制了并发的数量，防止系统过载。这里的400是并发的goroutine数量上限，可以根据实际情况调整。
线程安全 ：使用互斥锁mu确保对共享资源dataMap的访问是线程安全的，防止出现数据竞争问题。
*/
