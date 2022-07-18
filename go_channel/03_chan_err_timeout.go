package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go func() {
		err := execFunc()
		if err != nil {
			fmt.Println("execFunc err")
		}
		ch <- true
	}()
	select {
	case <-ch:
		fmt.Println("done")
		return
	case <-time.After(time.Second * 1):
		fmt.Println("TimeOut")
		return
	}
}

// mock TimeOut
func execFunc() error {
	time.Sleep(time.Second * 3)
}
