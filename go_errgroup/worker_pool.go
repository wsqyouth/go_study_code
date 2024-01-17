package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"

	"github.com/gammazero/workerpool"
)

const (
	workerpoolNum = 10
	maxNum        = 100
)

type Task struct {
	Name string
	Data []int
}

func processTask(task *Task) ([]int, error) {
	// 模拟处理任务
	if task.Name == "error" {
		return nil, errors.New("an error occurred")
	}
	return task.Data, nil
}

func main() {
	tasks := []*Task{
		{Name: "task1", Data: []int{11, 22, 33}},
		{Name: "task2", Data: []int{4, 5, 6}},
		{Name: "task3", Data: []int{7, 8, 9}},
	}

	wp := workerpool.New(workerpoolNum)
	wg := sync.WaitGroup{}
	errChan := make(chan error, len(tasks))
	results := make([][]int, len(tasks))

	for i, task := range tasks {
		wg.Add(1)
		index := i
		currentTask := task
		wp.Submit(func() {
			defer wg.Done()
			result, err := processTask(currentTask)
			if err != nil {
				errChan <- err
				return
			}
			results[index] = result
		})
	}

	wg.Wait()
	wp.StopWait()

	select {
	case err := <-errChan:
		fmt.Println("Error:", err)
	default:
		fmt.Println("All tasks completed successfully")
	}

	// 合并结果
	var allResults []int
	for _, result := range results {
		allResults = append(allResults, result...)
	}

	// 排序结果
	sort.Slice(allResults, func(i, j int) bool { return allResults[i] < allResults[j] })

	fmt.Println("Results:", allResults)
}

// 学习使用wokerpool的并发处理模型
