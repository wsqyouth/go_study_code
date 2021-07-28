package main

import (
	"fmt"
	"sync"
)

func count(a, b int, n *sync.WaitGroup) {
	c := a + b
	fmt.Printf("The Result of %d + %d = %d\n", a, b, c)
	defer n.Done() //goroutinue完成后, WaitGroup的计数-1
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1) // WaitGroup的计数加1
		go count(i, i+1, &wg)
	}
	wg.Wait() //等待所有goroutine执行完毕
}

//方式1： 使用sync包同步goroutine
// WaitGroup 等待一组goroutinue执行完毕.
// 主程序调用 Add 添加等待的goroutinue数量.
// 每个goroutinue在执行结束时调用 Done ，此时等待队列数量减1.
// 主程序通过Wait阻塞，直到等待队列为0.
