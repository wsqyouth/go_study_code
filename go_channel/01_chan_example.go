package main

import (
	"fmt"
	"time"
)

func main() {
	//readUnbufferedChanDemo()
	//readChanMultiWayDemo()
	readBufferedChannelDemo()
	fmt.Println("vim-go")
}

// 无缓冲通道,读数据总是阻塞的
func readUnbufferedChanDemo() {

	var ch1 chan bool
	ch1 = make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine filled data: ", i)
		}
		ch1 <- true
		fmt.Println("filled finished")
	}()

	<-ch1
	fmt.Println("main over")
}

// 关闭chan并读取
func readChanMultiWayDemo() {

	ch1 := make(chan int)

	go sendData(ch1)
	//1. 第一种方式读数据
	/*
		for {
			time.Sleep(1 * time.Second)
			v, ok := <-ch1
			if !ok {
				fmt.Println("read finished")
				break
			}
			fmt.Println(v)
		}
	*/
	//2. 第二种方式读数据
	time.Sleep(1 * time.Second)
	for v := range ch1 {
		fmt.Println("range read", v)
	}
	fmt.Println("main over")
}
func sendData(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	//close
	close(ch)
}

// 缓冲channel
func readBufferedChannelDemo() {

	/**
	 * 非缓冲 make(chan T)
	 * 缓冲 make(chan T, capacity)
	 * 发送： 缓冲区的数据满了，才会阻塞
	 * 接受： 缓冲区的数据空了，才会阻塞
	 */
	ch2 := make(chan int, 100)
	fmt.Println(len(ch2), cap(ch2))

	go sendData(ch2)

	for v := range ch2 {
		fmt.Println(len(ch2), cap(ch2))
		fmt.Println("range read", v)
	}
	fmt.Println("main over")
}

func opSingleChannel() {
	/**
	  双向通道：
	      chan T
	           chan <- data, 发送数据，写出
	           data <- chan, 获取数据，读取

	  单向通道： 定向
	      chan <- T, 只支持写
	      <-chan T, 只支持读
	*/

	ch3 := make(chan<- int) // 单项，只能写，不能读
	ch4 := make(<-chan int) // 单项。只能读，不能写

	fmt.Println(ch3)
	fmt.Println(ch4)
}
