package main

import (
	_ "expvar"
	"fmt"
	"math/rand"
	"time"
)

// 待处理的工作
type Job struct {
	ID      int
	Payload string
}

// worker进行消费处理
type Worker struct {
	ID         int
	JobQueue   chan Job
	QuitChan   chan bool
	WorkerPool chan chan Job
}

// 开启一个worker进行消费,同时将总的任务池记录下来
func NewWorker(id int, workerPool chan chan Job) Worker {
	worker := Worker{
		ID:         id,
		JobQueue:   make(chan Job),
		QuitChan:   make(chan bool),
		WorkerPool: workerPool,
	}
	return worker
}

func doWorkJob(workerID int, jobID int) {
	fmt.Printf("worker%d: started job %d\n", workerID, jobID)
	// Do the work (simulate by sleeping for a random amount of time).
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	// Print the results.
	fmt.Printf("worker%d: completed job %d\n", workerID, jobID)
}

func (w *Worker) Start() {
	go func() {
		for {
			// 将当前woker的工作注册到工作池,表示自己可以处理job
			w.WorkerPool <- w.JobQueue
			select {
			// 监听从JobQueue获取工作进行处理
			case job := <-w.JobQueue:
				// We have received a work request.
				doWorkJob(w.ID, job.ID)
			case <-w.QuitChan:
				// asked to stop.
				fmt.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

func (w Worker) stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Dispatcher A pool of workers channels that are registered with the dispatcher
type Dispatcher struct {
	WorkerPool chan chan Job
	JobQueue   chan Job
	MaxWorkers int
}

// NewDispatcher 指定有多少个工人在处理,每个工人最多能出来多少个job
func NewDispatcher(jobQueueSize, maxWorkers int) *Dispatcher {
	workerQueue := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: workerQueue,
		JobQueue:   make(chan Job, jobQueueSize),
		MaxWorkers: maxWorkers,
	}
}

// Run 多个workers启动处理
func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i+1, d.WorkerPool)
		worker.Start()
	}

	// 分发job
	go d.dispatch()

	// 进行监控
	go d.monitor()
}

func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue:
			// 监听JobQueue总队列
			go func() {
				// Try to obtain a worker job channel that is available.
				// This will block until a worker is idle
				jobChannel := <-d.WorkerPool
				// Dispatch the job to the worker job channel
				jobChannel <- job //这个woker的工作满了怎么办?
			}()
		}
	}
}

func (d *Dispatcher) monitor() {
	maxJobNum := cap(d.JobQueue)
	alertJobNum := maxJobNum / 4
	warningJobNum := maxJobNum / 4 * 3
	for {
		select {
		case <-time.After(1 * time.Second):
			currentJobNum := len(d.JobQueue)
			if currentJobNum >= warningJobNum {
				fmt.Printf("Warning--> currentJobNum%d: warningJobNum job %d\n", currentJobNum, warningJobNum)
			} else if currentJobNum >= alertJobNum {
				fmt.Printf("Alert-->currentJobNum%d: alertJobNum job %d\n", currentJobNum, alertJobNum)
			} else {
				fmt.Printf("Normal-->currentJobNum%d: maxJobNum job %d\n", currentJobNum, maxJobNum)
			}
		}
	}
}

func main() {
	// Create a new dispatcher with a job queue size of 100 and a maximum of 8 workers.
	dispatcher := NewDispatcher(100, 8)
	dispatcher.Run()
	// Add some jobs to the job queue.
	for i := 1; i <= 2000; i++ {
		job := Job{
			ID:      i,
			Payload: fmt.Sprintf("job %d", i),
		}
		dispatcher.JobQueue <- job
	}
	// Wait for the jobs to complete.
	time.Sleep(30 * time.Second)
}

/*
总结：
在上文的基础上添加了对每个worker工作队列的监控告警
后续优化: 这个dispatcher可以用于多个主题的任务,推而广之,发挥最大功效
参考:
https://gist.github.com/harlow/dbcd639cf8d396a2ab73
*/
