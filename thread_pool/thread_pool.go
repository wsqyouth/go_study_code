package main

import (
	"sync"
)

// 定义任务
type ThreadPool struct {
	poolSize    int  // 协程池大小
	termination bool // 终止标志

	queueSize     int              // 队列长度
	jobsQueue     chan interface{} // 任务队列
	resultsQueue  chan interface{} // 结果队列
	jobsClosed    bool             // 任务结束标志
	resultsClosed bool             // 结果结束标志
}

//
// @Description: 新建一个协程池
// @param poolSize 池大小
// @param queueSize 任务队列长度
// @return *ThreadPool 协程池对象
// @return error
//
func NewThreadPool(poolSize, queueSize int) (*ThreadPool, error) {
	if poolSize <= 0 {
		return nil, ErrInvalidPoolSize
	}
	if queueSize <= 0 {
		return nil, ErrInvalidQueueSize
	}

	tp := &ThreadPool{}
	tp.poolSize = poolSize
	tp.queueSize = queueSize
	tp.jobsQueue = make(chan interface{}, queueSize)
	tp.resultsQueue = make(chan interface{}, queueSize)
	tp.termination = false
	tp.jobsClosed = false
	tp.resultsClosed = false
	return tp, nil
}

//
// @Description: 协程池启动
// @receiver t
//
func (t *ThreadPool) Start() {
	// 另起一个线程
	go func() {
		var wg sync.WaitGroup
		for i := 0; i < t.poolSize; i++ {
			wg.Add(1)
			// 启动worker
			go t.NewWorker(&wg)
		}
		// 保证任务执行完成前不退出
		wg.Wait()
		t.CloseResultsQueue()
		t.termination = true
	}()
}

//
// @Description: 构造一个worker来监听任务队列
// @receiver t
// @param wg 用于协程之间同步
//
func (t *ThreadPool) NewWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case job, ok := <-t.jobsQueue:
			if !ok {
				return
			}
			// 根据job的类型来执行任务
			switch j := job.(type) {
			case *RunnableJob:
				j.Handler(j.Params...)
			case *CallableJob:
				r := j.Handler(j.Params...)
				t.resultsQueue <- r
			case Runnable:
				j.Run()
			case Callable:
				r := j.Call()
				t.resultsQueue <- r
			default:
				panic(ErrInvalidJobType)
			}
		}
	}
}

//
// @Description: 提交任务
// @receiver t
// @param handler 方法名
// @param params 方法参数
//
func (t *ThreadPool) SubmitJob(handler interface{}, params []interface{}) {
	switch h := handler.(type) {
	case RunnableFunc:
		job := &RunnableJob{
			Handler: h,
			Params:  params,
		}
		t.jobsQueue <- job
	case CallableFunc:
		job := &CallableJob{
			Handler: h,
			Params:  params,
		}
		t.jobsQueue <- job
	default:
		panic(ErrInvalidJobType)
	}
}

//
// @Description: 提交任务
// @receiver t
// @param job 任务对象
//
func (t *ThreadPool) Submit(job interface{}) {
	t.jobsQueue <- job
}

//
// @Description: 获取任务结果并且回传给上层
// @receiver t
// @param handler
//
func (t *ThreadPool) GetResults(handler func(interface{})) {
	go func() {
		for result := range t.resultsQueue {
			handler(result)
		}
	}()
}

//
// @Description: 获取任务结果并且回传给上层
// @receiver t
// @param job
//
//func (t *ThreadPool) Get(job Callable) {
//	go func() {
//		for result := range t.resultsQueue {
//			job.Back(result)
//		}
//	}()
//}

//
// @Description: 关闭结果队列
// @receiver t
//
func (t *ThreadPool) CloseResultsQueue() {
	if !t.resultsClosed {
		close(t.resultsQueue)
	}
	t.resultsClosed = true
}

//
// @Description: 关闭任务队列
// @receiver t
//
func (t *ThreadPool) CloseJobsQueue() {
	if !t.jobsClosed {
		close(t.jobsQueue)
	}
	t.jobsClosed = true
}

//
// @Description: 关闭协程池
// @receiver t
//
func (t *ThreadPool) Close() {
	t.CloseResultsQueue()
	t.CloseJobsQueue()
	t.termination = true
}

//
// @Description: 等待所有任务执行完毕，会阻塞主进程
// @receiver t
//
func (t *ThreadPool) Wait() {
	t.CloseJobsQueue()
	// 等待所有任务执行完毕
	for {
		if t.IsDone() {
			break
		}
	}
}

//
// @Description: 判断协程池中所有任务是否执行完成
// @receiver t
// @return bool
//
func (t *ThreadPool) IsDone() bool {
	return t.termination
}
