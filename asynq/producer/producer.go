package main

import (
	"log"
	"time"

	"github.com/hibiken/asynq"
	// "tasks"
	"github.com/wsqyouth/coopers_go_code/asynq/tasks"
)

const redisAddr = "127.0.0.1:6379"

func main() {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: redisAddr})
	defer client.Close()

	// ------------------------------------------------------
	// Example 1: Enqueue task to be processed immediately. Use (*Client).Enqueue method.
	//         立即将任务放入队列,使用Enqueue方法
	// ------------------------------------------------------

	task, err := tasks.NewEmailDeliveryTask(42, "some:template:id")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ------------------------------------------------------------
	// Example 2: Schedule task to be processed in the future.Use ProcessIn or ProcessAt option.
	//            任务稍后处理,使用ProcessIn方法
	// ------------------------------------------------------------

	info, err = client.Enqueue(task, asynq.ProcessIn(1*time.Minute))
	if err != nil {
		log.Fatalf("could not schedule task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ----------------------------------------------------------------------------
	// Example 3: Set other options to tune task processing behavior. Options include MaxRetry, Queue, Timeout, Deadline, Unique etc.
	//            针对特定taskType,使用option函数设置处理行为，比如最大重试次数，超时等参数
	// ----------------------------------------------------------------------------

	client.SetDefaultOptions(tasks.TypeImageResize, asynq.MaxRetry(10), asynq.Timeout(3*time.Minute))

	task, err = tasks.NewImageResizeTask("https://baidu.com/image.jpg")
	if err != nil {
		log.Fatalf("could not create task: %v", err)
	}
	info, err = client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	// ---------------------------------------------------------------------------
	// Example 4: Pass options to tune task processing behavior at enqueue time.Options passed at enqueue time override default ones.
	//            在Enqueue时的参数设置可以覆盖全局默认设置
	// ---------------------------------------------------------------------------

	info, err = client.Enqueue(task, asynq.Queue("critical"), asynq.Timeout(30*time.Second))
	if err != nil {
		log.Fatal("could not enqueue task: %v", err)
	}
	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)
}
