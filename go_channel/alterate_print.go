package main

import (
	"fmt"
	"time"
)

type Token struct{}

func main() {
	fmt.Println("vim-go")
	AlternatelyPrint(5)
}

func AlternatelyPrint(total int) {
	// 初始化多个channel，用于goroutine之间通信
	channels := make([]chan Token, total)
	for i := 0; i < total; i++ {
		channels[i] = make(chan Token)
	}

	for i := 0; i < total; i++ {
		go func(index int, current chan Token, nextChan chan Token) {
			for {
				<-current
				fmt.Printf("Goroutine %d\n", index)
				time.Sleep(time.Second)
				nextChan <- Token{}
			}
		}(i+1, channels[i], channels[(i+1)%total])
	}
	channels[0] <- Token{}
	select {}
}
