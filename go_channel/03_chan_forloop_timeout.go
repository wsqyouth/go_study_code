package main

import (
	"fmt"
	"time"
)

func main() {
	news := make(chan string)
	go newsFeed(news)
	printAllNews(news)
	fmt.Println("done")
}

func newsFeed(ch chan string) {
	for i := 0; i < 2; i++ {
		time.Sleep(time.Millisecond * 400)
		ch <- fmt.Sprintf("News: %d", i+1)
	}
}
func printAllNews(news chan string) {
	for {
		select {
		case msg := <-news:
			fmt.Println(msg)
		case <-time.After(time.Second * 1):
			fmt.Println("TimeOut: News feed finished")
			return
		}
	}
}
