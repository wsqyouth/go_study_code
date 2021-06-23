package main

import (
	"errors"
	"fmt"

	"golang.org/x/sync/errgroup"
)

//errgroup包在sync.WaitGroup功能的基础上，增加了错误传递，以及在发生不可恢复的错误时取消整个goroutine集合，或者等待超时
func main() {

	var g errgroup.Group

	g.Go(func() error {
		fmt.Println("Goroutine 1")
		return nil //返回nil
	})
	g.Go(func() error {
		fmt.Println("Goroutine 2")
		return errors.New("Goroutine error") //返回自定义错误
	})
	if err := g.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Success")
	}

}
