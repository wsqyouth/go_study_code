package main

import (
	"fmt"
	"time"
)

func count(a, b int, exitChan chan bool) {
	c := a + b
	fmt.Printf("The Result of %d + %d = %d\n", a, b, c)
	time.Sleep(time.Second * 2)
	exitChan <- true
}

func main() {
	exitChan := make(chan bool, 10) //声明并分配管道内存
	for i := 0; i < 10; i++ {
		go count(i, i+1, exitChan)
	}

	for i := 0; i < 10; i++ {
		<-exitChan //取信号数据，如果取不到则阻塞
	}
	close(exitChan)
}

//方式2：通过channel实现goroutine之间的同步:

//通过channel能在多个groutine之间通讯，当一个goroutine完成时候向channel发送退出信号,
//等所有goroutine退出时候，利用for循环channe去channel中的信号，若取不到数据会阻塞原理，等待所有goroutine执行完毕
//使用该方法有个前提是你已经知道了你启动了多少个goroutine

//ref: https://github.com/KeKe-Li/For-learning-Go-Tutorial/blob/master/src/chapter10/01.0.md
