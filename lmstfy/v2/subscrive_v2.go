package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Message 定义
type Message struct {
	ID      int
	Header  string
	Data    string
	JobType string
}

// Job 处理结果
type JobResult struct {
	Success bool
	Message string
	Error   error
}

// JobProcessor 定义业务处理接口
type JobProcessor interface {
	Process(msg Message) JobResult
}

// 具体业务处理器
type Queue1Processor struct{}
type Queue2Processor struct{}

func (p *Queue1Processor) Process(msg Message) JobResult {
	// 模拟队列1的业务处理
	time.Sleep(100 * time.Millisecond)
	return JobResult{
		Success: true,
		Message: fmt.Sprintf("Queue1 processed msg: %d", msg.ID),
	}
}

func (p *Queue2Processor) Process(msg Message) JobResult {
	// 模拟队列2的业务处理
	time.Sleep(150 * time.Millisecond)
	if msg.ID%3 == 0 {
		return JobResult{
			Success: false,
			Error:   fmt.Errorf("queue2 failed to process msg: %d", msg.ID),
		}
	}
	return JobResult{
		Success: true,
		Message: fmt.Sprintf("Queue2 processed msg: %d", msg.ID),
	}
}

// Queue 定义
type Queue struct {
	name string
	ch   chan Message
}

func NewQueue(name string, size int) *Queue {
	return &Queue{
		name: name,
		ch:   make(chan Message, size),
	}
}

// Processor 定义
type Processor struct {
	name      string
	msgChan   chan Message
	processor JobProcessor
	threads   int
}

func NewProcessor(name string, bufferSize int, threads int, processor JobProcessor) *Processor {
	return &Processor{
		name:      name,
		msgChan:   make(chan Message, bufferSize),
		processor: processor,
		threads:   threads,
	}
}

// Processor methods
func (p *Processor) Input() chan<- Message {
	return p.msgChan
}

func (p *Processor) Start(ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < p.threads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			p.processLoop(ctx, id)
		}(i)
	}
}

func (p *Processor) processLoop(ctx context.Context, threadID int) {
	for {
		select {
		case <-ctx.Done():
			// 处理完当前channel中的所有消息后退出
			for {
				select {
				case msg, ok := <-p.msgChan:
					if !ok {
						return
					}
					p.processMessage(msg, threadID)
				default:
					return
				}
			}
		case msg, ok := <-p.msgChan:
			if !ok {
				return
			}
			p.processMessage(msg, threadID)
		}
	}
}

func (p *Processor) processMessage(msg Message, threadID int) {
	result := p.processor.Process(msg)
	if result.Success {
		fmt.Printf("[Processor-%s-Thread-%d] Successfully processed message %d: %s\n",
			p.name, threadID, msg.ID, result.Message)
	} else {
		fmt.Printf("[Processor-%s-Thread-%d] Failed to process message %d: %s\n",
			p.name, threadID, msg.ID, result.Error)
	}
}

// Subscriber 定义
type Subscriber struct {
	name       string
	queue      *Queue
	processor  *Processor
	subThreads int
}

func NewSubscriber(name string, queue *Queue, processor *Processor, subThreads int) *Subscriber {
	return &Subscriber{
		name:       name,
		queue:      queue,
		processor:  processor,
		subThreads: subThreads,
	}
}

// Subscriber methods
func (s *Subscriber) Start(ctx context.Context, wg *sync.WaitGroup) {
	// 启动processor
	s.processor.Start(ctx, wg)

	// 启动subscriber goroutines
	for i := 0; i < s.subThreads; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			s.subscribeLoop(ctx, id)
		}(i)
	}
}

func (s *Subscriber) subscribeLoop(ctx context.Context, threadID int) {
	for {
		select {
		case <-ctx.Done():
			// 优雅退出：确保队列中的消息都被消费
			for {
				select {
				case msg, ok := <-s.queue.ch:
					if !ok {
						return
					}
					fmt.Printf("[Subscriber-%s-Thread-%d] Processing remaining message: %d\n",
						s.name, threadID, msg.ID)
					s.processor.Input() <- msg
				default:
					return
				}
			}
		case msg := <-s.queue.ch:
			fmt.Printf("[Subscriber-%s-Thread-%d] Received message: %d\n",
				s.name, threadID, msg.ID)
			s.processor.Input() <- msg
		}
	}
}

// Job生成器
func genJob(queueName string, id int) Message {
	return Message{
		ID:      id,
		Header:  queueName,
		Data:    fmt.Sprintf("data-%d", id),
		JobType: fmt.Sprintf("%s-job", queueName),
	}
}

func main() {
	// 创建队列
	queue1 := NewQueue("queue1", 100)
	queue2 := NewQueue("queue2", 100)

	// 创建processor
	proc1 := NewProcessor("proc1", 100, 2, &Queue1Processor{})
	proc2 := NewProcessor("proc2", 100, 3, &Queue2Processor{})

	// 创建subscriber
	sub1 := NewSubscriber("sub1", queue1, proc1, 3)
	sub2 := NewSubscriber("sub2", queue2, proc2, 2)

	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	// 启动subscribers
	sub1.Start(ctx, &wg)
	sub2.Start(ctx, &wg)

	// 生成测试数据
	go func() {
		for i := 0; i < 10; i++ {
			queue1.ch <- genJob("queue1", i)
			queue2.ch <- genJob("queue2", i)
			time.Sleep(200 * time.Millisecond)
		}

		// 等待一段时间后发送退出信号
		time.Sleep(2 * time.Second)
		fmt.Println("Initiating graceful shutdown...")
		cancel()
	}()

	// 等待所有goroutine完成
	wg.Wait()
	fmt.Println("All processors and subscribers have completed")
}
