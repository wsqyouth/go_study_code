package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/groupcache/singleflight"
)

func NewDelayReturn(dur time.Duration, n int) func() (interface{}, error) {
	return func() (interface{}, error) {
		time.Sleep(dur)
		return n, nil

	}

}

func main() {
	g := singleflight.Group{}
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		ret, err := g.Do("key", NewDelayReturn(time.Second*1, 1))
		if err != nil {
			panic(err)

		}
		fmt.Printf("key-1 get %v\n", ret)
		wg.Done()

	}()
	go func() {
		time.Sleep(100 * time.Millisecond) // make sure this is call is later
		ret, err := g.Do("key", NewDelayReturn(time.Second*2, 2))
		if err != nil {
			panic(err)

		}
		fmt.Printf("key-2 get %v\n", ret)
		wg.Done()

	}()
	go func() {
		time.Sleep(100 * time.Millisecond) // make sure this is call is later
		ret, err := g.Do("key_other", NewDelayReturn(time.Second*10, 10))
		if err != nil {
			panic(err)

		}
		fmt.Printf("key-other get %v\n", ret)
		wg.Done()

	}()
	wg.Wait()

}
