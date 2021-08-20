package main

/*
模拟共享数据库连接
*/
import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines   = 5 //创建goroutine数量
	pooledResources = 2 //池中资源数量
)

// dbConnection 模拟要共享的资源
type dbConnection struct {
	ID int32
}

// 实现接口，用来完成资源释放
func (dbConn *dbConnection) Close() error {
	log.Println("Close: connection", dbConn.ID)
	return nil
}

var idCounter int32

// 工厂函数，需要新连接时调用
func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: new connection", id)

	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// 创建管理的连接池
	p, err := New(createConnection, pooledResources)
	if err != nil {
		log.Println(err)
		panic("New error")
	}

	// 使用池中的连接完成查询
	for query := 0; query < maxGoroutines; query++ {
		//每个goroutine要复制的queryID,否则所有查询共享同一个查询变量
		go func(q int) {
			performQueries(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()

	log.Println("Shutdown Program.")
	p.Close()
	fmt.Println("Done")
}

// 测试连接的资源池
func performQueries(query int, p *Pool) {
	//从池中请求连接
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	//释放
	defer p.Release(conn)

	//用例模拟查询耗时
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("uid[%d] connectID[%d]\n", query, conn.(*dbConnection).ID)
}
