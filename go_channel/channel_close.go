package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.NumGoroutine())
	sendTasksCheckClose()
	time.Sleep(time.Second)
	runtime.GC()
	fmt.Println(runtime.NumGoroutine())
}

func doCheckClose(taskCh chan int) {
	for {
		select {
		case t, ok := <-taskCh:
			if !ok {
				fmt.Println("taskCh has been closed")
				return
			}
			time.Sleep(time.Millisecond)
			fmt.Printf("task %d is done\n", t)
		}
	}
}

func sendTasksCheckClose() {
	taskCh := make(chan int, 10)
	go doCheckClose(taskCh)
	for i := 0; i < 20; i++ {
		taskCh <- i
	}
	close(taskCh)
}

/*
这个代码主要是学习如何正确的关闭channel, 如何在判断chanel关闭时退出，否则一直阻塞会导致goroutin泄漏
*/
