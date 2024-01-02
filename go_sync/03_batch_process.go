package main

import (
	"fmt"
	"sync"
)

func main() {
	tasks := []int{1, 2, 3, 4, 5}
	results := ProcessTasks(tasks)

	count := 0
	for v := range results {
		count += v
	}

	fmt.Printf("Total processed tasks: %d\n", count)
}

func ProcessTasks(tasks []int) <-chan int {
	ch := make(chan int, 0)
	var wg sync.WaitGroup
	n := len(tasks)
	wg.Add(n)
	for i := 0; i < n; i++ {
		i, v := i, tasks[i]
		go func(i, v int) {
			defer wg.Done()
			//fmt.Println(i, v)
			ch <- v
		}(i, v)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

/*
这里我实现的时候,ProcessTasks没有使用go func() {}  直接:
wg.Wait()
close(ch)
这样是错误的,原因在于:
如果你不在新的goroutine中调用wg.Wait()，那么ProcessTasks函数会阻塞在wg.Wait()，并且永远不会有机会关闭通道。
这就导致了主函数中的for循环永远在等待新的输入，但是因为通道没有关闭，所以新的输入永远不会到来，这就是死锁。
通过在新的goroutine中调用wg.Wait()，你可以让ProcessTasks函数立即返回通道，然后在所有任务完成后关闭通道。这样，主函数就可以开始从通道中读取数据，而不会被阻塞。
*/
