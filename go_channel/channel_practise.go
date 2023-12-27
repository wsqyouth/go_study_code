package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	panicExample()
	fmt.Println("hello")
}

/*
总结复盘:
学习资料:https://go.dev/doc/effective_go#channels
*/
// 问题1: 如何交替打印数字
func altePrint() {
	ch1 := make(chan struct{}, 0)
	ch2 := make(chan struct{}, 0)
	printA := func() {
		for {
			<-ch1
			fmt.Println("aaa")
			time.Sleep(1 * time.Second)
			ch2 <- struct{}{}
		}

	}

	printB := func() {
		for {
			<-ch2
			fmt.Println("bbb")
			time.Sleep(1 * time.Second)
			ch1 <- struct{}{}
		}
	}
	go printA()
	go printB()

	go func() {
		ch1 <- struct{}{}
	}()
	time.Sleep(time.Hour * 1)
}

func alterPrintWithSelect() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	printAB := func() {
		for {
			select {
			case <-ch1:
				fmt.Println("aaaa")
				time.Sleep(1 * time.Second)
				ch2 <- struct{}{}
			case <-ch2:
				fmt.Println("bbbb")
				time.Sleep(1 * time.Second)
				ch1 <- struct{}{}
			}
		}
	}
	go printAB()
	go printAB() // 这里还需要一个呢

	go func() {
		ch1 <- struct{}{}
	}()
	time.Sleep(1 * time.Hour)
}

// 问题2: 如何并发处理数据, 学习wg的用法,但是有问题,共享变量自增
func processCon() {
	var cnt int
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cnt++
		}()
	}
	wg.Wait()
	fmt.Println("100次自增", cnt)
}

func processConV1() {
	var cnt int64 // 定义为普通类型,但是会用原子操作
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&cnt, 1)
		}()
	}
	wg.Wait()
	fmt.Println("100次自增", cnt)
}

func processConV2() {
	var cnt int64 // 定义为普通类型
	var lock sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			cnt++
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("100次自增", cnt)
}

// 问题3: 使用chan实现一个生产者消费者模型
func product(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) //fatal error: all goroutines are asleep - deadlock!
}
func consumer(ch <-chan int, done chan<- struct{}) {
	for v := range ch {
		fmt.Println(v)
	}
	done <- struct{}{}
}

func productConsumer() {
	ch := make(chan int, 5)
	done := make(chan struct{})
	go product(ch)
	go consumer(ch, done)
	<-done
}

// 问题4: 如何使用wg并发消费的消费者模型
func processInput() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -1} // 这是我们的流水数据
	ch := make(chan int)
	fmt.Println("Number of goroutines:", runtime.NumGoroutine())
	maxGoroutine := len(numbers)
	if maxGoroutine > runtime.NumGoroutine() {
		maxGoroutine = runtime.NumGoroutine()
	}
	var wg sync.WaitGroup
	for i := 0; i < maxGoroutine; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for num := range ch {
				fmt.Println("process num:", num)
			}
		}()
	}
	for _, num := range numbers {
		ch <- num
	}
	close(ch) // 一定要放到wait之前
	wg.Wait() // 正确的顺序是先关闭 channel，然后等待所有的 goroutine 结束。
	fmt.Println("done")
}

// 问题5: 如何让chan进行panic
func panicExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("panic recover: ", r)
		}
		fmt.Println("defer")
	}()
	// 场景1：关闭nil的chan会panic
	//var ch chan int
	//close(ch)

	// 场景2: 关闭已经关闭的ch会panic
	//ch2 := make(chan struct{})
	//close(ch2)
	//close(ch2)

	// 场景3: 向一个关闭的ch写会panic
	ch3 := make(chan struct{}, 5)
	close(ch3)
	ch3 <- struct{}{}
}
