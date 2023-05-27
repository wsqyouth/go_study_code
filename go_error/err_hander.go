package main

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	handlers := []func() error{
		func() error {
			fmt.Println("func 1 run in start")
			time.Sleep(3 * time.Microsecond)
			a := 1
			b := 0
			fmt.Println(a / b)
			return nil
		},
		func() error {
			fmt.Println("func2 run in start")
			time.Sleep(3 * time.Microsecond)
			fmt.Println("func2 run in end")
			return nil
		},
	}
	if err := processHandlers(handlers...); err != nil {
		fmt.Println(err)
	}
	fmt.Println("Hello World")
}

func processHandlers(handlers ...func() error) (err error) {
	var wg sync.WaitGroup
	var once sync.Once
	for i := range handlers {
		wg.Add(1)
		go func(handler func() error) {
			defer func() {
				if e := recover(); e != nil {
					buf := make([]byte, 2048)
					buf = buf[:runtime.Stack(buf, false)]
					fmt.Printf("[PANIC]%v\n%s\n", e, buf)
					once.Do(func() {
						err = errors.New("panic found in call handlers")
					})
				}
				wg.Done()
			}()
			if e := handler(); e != nil {
				once.Do(func() {
					err = e
				})
			}
		}(handlers[i])
	}
	wg.Wait()
	return err
}
