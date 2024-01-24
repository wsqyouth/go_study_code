package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	tasks := make([]Task, len(nums))
	results := make([]int, len(nums))

	for i, num := range nums {
		tasks[i] = &SquareTask{
			Num:    num,
			Result: &results[i],
		}
	}

	ctx := context.Background()
	if err := concurrentExec(ctx, tasks, 2); err != nil {
		log.Fatalf("Failed to execute tasks: %v", err)
	}

	fmt.Println(results) // Output: [1 4 9 16 25]
}

type SquareTask struct {
	Num    int
	Result *int
}

func (t *SquareTask) Execute(ctx context.Context) error {
	*t.Result = t.Num * t.Num
	return nil
}

type Task interface {
	Execute(ctx context.Context) error
}

func concurrentExec(ctx context.Context, tasks []Task, maxGoroutineNum int) error {
	if len(tasks) == 0 {
		return nil
	}

	if len(tasks) < maxGoroutineNum {
		maxGoroutineNum = len(tasks)
	}

	ch := make(chan Task)
	g, ctx := errgroup.WithContext(ctx)

	for i := 0; i < maxGoroutineNum; i++ {
		g.Go(func() error {
			for task := range ch {
				if err := task.Execute(ctx); err != nil {
					return err
				}
			}
			return nil
		})
	}

	for _, task := range tasks {
		ch <- task
	}
	close(ch)

	return g.Wait()
}

/*
我们定义了一个Task接口，它包含一个Execute方法。这样，我们可以为不同类型的任务实现这个接口，然后将它们传递给concurrentExec函数。
这样，我们就可以在不修改concurrentExec函数的情况下，处理不同类型的任务，从而遵循开闭原则。

这个设计更加通用，可以应用于各种类型的任务。同时，它也解决了任务类型耦合的问题，使得concurrentExec函数可以处理任何实现了Task接口的任务。
*/
