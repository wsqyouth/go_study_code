package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Service interface {
	Close(chan struct{})
}

type MyService struct {
	name string
}

func (s *MyService) Close(c chan struct{}) {
	// 模拟关闭服务所需的时间
	time.Sleep(time.Duration(1+rand.Intn(5)) * time.Second)
	fmt.Println("close name:", s.name)
	c <- struct{}{}
}

func main() {
	services := []Service{
		&MyService{name: "Service 1"},
		&MyService{name: "Service 2"},
		&MyService{name: "Service 3"},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var wg sync.WaitGroup
	for _, service := range services {
		wg.Add(1)
		go func(srv Service) {
			defer wg.Done()

			c := make(chan struct{}, 1)
			go srv.Close(c)

			select {
			case <-c:
			case <-ctx.Done():
			}
		}(service)
	}

	wg.Wait()
	fmt.Println("All services closed or timed out")
}

/*
在给定的超时时间内关闭一组服务。它使用了 context，sync.WaitGroup 和 goroutines 来实现这个功能
*/
