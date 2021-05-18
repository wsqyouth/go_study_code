package main

import (
	"fmt"

	"time"
)

func main() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("monitor stop...")
				return
			default:
				fmt.Println("monitor continuing...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	//通知停止监控
	stop <- true

	//继续观察看是否go routine是否仍在执行,如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)
}

//文档参考：
// https://github.com/iswbm/GolangCodingTime/blob/master/source/c04/c04_08.md
