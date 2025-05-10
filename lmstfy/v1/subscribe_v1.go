package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 消息结构
type Message struct {
	ID     int
	Header string
	Data   string
}

// 模拟队列
type Queue struct {
	name string
	ch   chan Message
}

// Subscriber 结构
type Subscriber struct {
	name          string
	queue         *Queue
	processorChan chan Message
	subThreads    int
	procThreads   int
}

// Processor 结构
type Processor struct {
	msgChan chan Message
}

// NewQueue 创建队列
func NewQueue(name string, size int) *Queue {
	return &Queue{
		name: name,
		ch:   make(chan Message, size),
	}
}

// NewSubscriber 创建订阅者
func NewSubscriber(name string, queue *Queue, subThreads, procThreads int) *Subscriber {
	return &Subscriber{
		name:          name,
		queue:         queue,
		processorChan: make(chan Message, 100),
		subThreads:    subThreads,
		procThreads:   procThreads,
	}
}

// Start 启动订阅者
func (s *Subscriber) Start(ctx context.Context) {
	// 启动subscriber goroutines
	var wg sync.WaitGroup

	// 启动N个消费队列的goroutine
	for i := 0; i < s.subThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			s.subscribe(ctx, id)
		}(i)
	}

	// 启动M个处理消息的goroutine
	for i := 0; i < s.procThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			s.process(ctx, id)
		}(i)
	}

	wg.Wait()
}

// subscribe 消费队列消息写入处理channel
func (s *Subscriber) subscribe(ctx context.Context, threadID int) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.queue.ch:
			fmt.Printf("[Subscriber-%s-Thread-%d] Received message: %+v\n",
				s.name, threadID, msg)
			s.processorChan <- msg
		}
	}
}

// process 处理channel中的消息
func (s *Subscriber) process(ctx context.Context, threadID int) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.processorChan:
			// 模拟处理时间
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("[Processor-%s-Thread-%d] Processed and ACK message: %+v\n",
				s.name, threadID, msg)
		}
	}
}

func main() {
	// 创建两个队列
	queue1 := NewQueue("queue1", 100)
	queue2 := NewQueue("queue2", 100)

	// 创建两个subscriber，配置不同的并发数
	sub1 := NewSubscriber("sub1", queue1, 3, 2) // 3个订阅线程，2个处理线程
	sub2 := NewSubscriber("sub2", queue2, 2, 3) // 2个订阅线程，3个处理线程

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动subscribers
	go sub1.Start(ctx)
	go sub2.Start(ctx)

	// 向队列发送消息
	go func() {
		for i := 0; i < 5; i++ {
			queue1.ch <- Message{ID: i, Header: "queue1", Data: fmt.Sprintf("data-%d", i)}
			queue2.ch <- Message{ID: i, Header: "queue2", Data: fmt.Sprintf("data-%d", i)}
			time.Sleep(200 * time.Millisecond)
		}
		time.Sleep(2 * time.Second)
		cancel() // 2秒后停止所有处理
	}()

	<-ctx.Done()
	fmt.Println("Demo finished")
}
