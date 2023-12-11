package main

import (
	"context"
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

// 假设我们有一个map，键和值都是字符串
var myMap = map[string]string{
	"key1": "value1",
	"key2": "value2",
	"key3": "value3",
}

// 这是我们要并发执行的函数，它接受一个上下文，一个键和一个值
func processKeyValue(ctx context.Context, key, value string) error {
	// 在这里执行你的逻辑
	fmt.Printf("Processing %s: %s\n", key, value)
	return nil
}

func main() {
	g, ctx := errgroup.WithContext(context.Background())

	for key, value := range myMap {
		// 这里我们创建了key和value的副本，以便在并发的goroutine中安全地使用它们
		key := key
		value := value

		g.Go(func() error {
			if err := processKeyValue(ctx, key, value); err != nil {
				return errors.New(fmt.Sprintf("failed to process %s: %s", key, value))
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Encountered error: %v", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}

// 我们并发地处理一个map的每个键值对。我们使用errgroup包来并发执行任务，并在任何任务失败时取消所有任务。
