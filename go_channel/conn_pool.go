package main

import (
	"fmt"
	"sync"
)

// CustomClient 自定义客户端结构体
type CustomClient struct {
	ID string
}

// ClientChan 客户端管道
type ClientChan chan *CustomClient

// ConnPool 基于channel实现链接池管理
type ConnPool struct {
	sync.RWMutex
	poolSize      int
	clientChanMap map[string]ClientChan
}

// NewConnPool 创建一个新的连接池
func NewConnPool(poolSize int) *ConnPool {
	return &ConnPool{
		poolSize:      poolSize,
		clientChanMap: make(map[string]ClientChan),
	}
}

// CheckPoolSizeAndCapacity 检查poolSize和各个客户端池的容量
func (p *ConnPool) CheckPoolSizeAndCapacity() {
	p.RLock()
	defer p.RUnlock()

	fmt.Printf("Pool size: %d\n", p.poolSize)

	for addr, clientChan := range p.clientChanMap {
		fmt.Printf("Address: %s, Channel capacity: %d, Current length: %d\n", addr, cap(clientChan), len(clientChan))
	}
}

func main() {
	pool := NewConnPool(10)

	// 假设我们有一些地址和对应的客户端
	addrs := []string{"addr1", "addr2", "addr3"}
	for _, addr := range addrs {
		pool.clientChanMap[addr] = make(ClientChan, pool.poolSize)
	}

	// 检查poolSize和各个客户端池的容量
	pool.CheckPoolSizeAndCapacity()
}

/*
基于channel的通用连接池的通用实现:
这个代码比较好的实现了连接池,内部通过map+读写锁实现
比较特殊的点在于,整体map的容量是不受限的,但是每个地址对应的客户端管道的容量是固定的
*/
