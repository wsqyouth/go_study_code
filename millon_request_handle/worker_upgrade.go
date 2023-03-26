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
			// 将当前woker的工作注册到工作池
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
	go d.dispatch()
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

func main() {
	// Create a new dispatcher with a job queue size of 100 and a maximum of 8 workers.
	dispatcher := NewDispatcher(100, 8)
	dispatcher.Run()
	// Add some jobs to the job queue.
	for i := 1; i <= 20; i++ {
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
使用chatGpt升级了代码,使用woker进行消费处理,能够达到百万并发处理
这里的参数设置很巧妙,另外当woker数量满时,jobQueue满时都是需要考虑的
每开启一个woker时将当前woker放入chan,同时监听获取一个可用jobQueue是非常巧妙的
参考:
http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/
https://gist.github.com/harlow/dbcd639cf8d396a2ab73
*/
