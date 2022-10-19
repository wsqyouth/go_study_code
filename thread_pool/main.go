package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	TestCallableJob()
}

//
// @Description: 单线程 使用函数 求和
// @param t
//
func TestCallableJob() {
	for i := 100; i > 0; i-- {
		result := CallJob([]interface{}{i}...)
		fmt.Printf("result: %v\n", result)
	}
}

//===============================================================================================================//

//
// @Description: 多线程 使用函数 求和
// @param t
//
func TestThreadPoolCallableJob() {
	pool, err := NewThreadPool(10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	pool.GetResults(PrintfResult)
	pool.Start()

	for i := 100; i > 0; i-- {
		params := []interface{}{i}
		pool.SubmitJob(CallableFunc(CallJob), params)
	}
	pool.Wait()
	pool.Close() // 可以省略
}

//
// @Description: 任务函数
// @param v
// @return interface{}
//
func CallJob(v ...interface{}) interface{} {
	sum := 0
	number := v[0].(int)
	for i := 0; i <= number; i++ {
		sum += i
		time.Sleep(time.Millisecond * time.Duration(i))
	}
	fmt.Printf("number: %d\n", number)
	return sum
}

//
// @Description: 打印结果
// @param result
//
func PrintfResult(result interface{}) {
	fmt.Printf("result: %v\n", result)
}

//===============================================================================================================//

//
//  Task
//  @Description: 任务结构体
//
type Task struct {
	Number int
}

//
// @Description: Call方法的实现，满足Callable接口
// @receiver t
// @return interface{}
//
func (t *Task) Call() interface{} {
	sum := 0
	number := t.Number
	for i := 0; i <= number; i++ {
		sum += i
		time.Sleep(time.Millisecond * time.Duration(i))
	}
	return sum
}

//
// @Description: 多线程 使用接口 求和
// @param t
//
func TestThreadPoolCallable() {
	pool, err := NewThreadPool(10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	pool.GetResults(PrintfResult)
	pool.Start()

	for i := 100; i > 0; i-- {
		task := &Task{
			Number: i,
		}
		pool.Submit(task)
	}
	pool.Wait()
}

//===============================================================================================================//

//
// @Description: 任务函数 不带返回参数
// @param v
//
func RunJob(v ...interface{}) {
	sum := 0
	number := v[0].(int)
	for i := 0; i <= number; i++ {
		sum += i
		time.Sleep(time.Millisecond * time.Duration(i))
	}
	fmt.Printf("number: %d\n", number)
	fmt.Printf("sum: %d\n", sum)
}

//
// @Description: 多线程 使用函数 没有返回
// @param t
//
func TestThreadPoolRunnableJob() {
	pool, err := NewThreadPool(10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	pool.Start()

	for i := 100; i > 0; i-- {
		params := []interface{}{i}
		pool.SubmitJob(RunnableFunc(RunJob), params)
	}
	pool.Wait()
}

//===============================================================================================================//

//
// @Description: Run()方法实现，满足Runnable接口
// @receiver t
//
func (t *Task) Run() {
	sum := 0
	number := t.Number
	for i := 0; i <= number; i++ {
		sum += i
		time.Sleep(time.Millisecond * time.Duration(i))
	}
	fmt.Printf("number: %d\n", number)
	fmt.Printf("sum: %d\n", sum)
}

//
// @Description: 多线程 使用接口 没有返回值
// @param t
//
func TestThreadPoolRunnable() {
	pool, err := NewThreadPool(10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	pool.Start()

	for i := 100; i > 0; i-- {
		task := &Task{
			Number: i,
		}
		pool.Submit(task)
	}
	pool.Wait()
}

//===============================================================================================================//

type TaskCount struct {
	Count  *int32
	Number int
}

//
// @Description: 累加任务
// @receiver tc
//
func (tc *TaskCount) Run() {
	sum := 0
	number := tc.Number
	for i := 0; i <= number; i++ {
		sum += i
	}
	atomic.AddInt32(tc.Count, 1) // 原子加
	//*tc.Count += 1  // 不用原子加会出现错误
}

//
// @Description: 多线程累加全局变量
// @param t
//
func TestThreadPoolCount() {
	pool, err := NewThreadPool(10, 10)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	pool.Start()
	var globalCount int32 = 0

	for i := 1000000; i > 0; i-- {
		task := &TaskCount{
			Number: i,
			Count:  &globalCount,
		}
		pool.Submit(task)
	}
	pool.Wait()
	fmt.Println(globalCount)
}
