package main

import (
	"fmt"
	"sync"
)

func main() {
	goroutineNum := 4
	var wg sync.WaitGroup
	var mu sync.Mutex
	var resource int = 1
	for j := 0; j < goroutineNum; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			fmt.Println(resource)
			resource++
		}()
	}
	wg.Wait()
	fmt.Println("final:", resource)
}
