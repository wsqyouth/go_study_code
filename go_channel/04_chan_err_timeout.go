package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	ch := make(chan bool)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(fmt.Sprintf("recover: %#v", err))
			}
		}()
		//err := execFuncError()
		err := execFuncPanic()
		if err != nil {
			fmt.Println("execFunc err")
		}
		ch <- true
	}()
	select {
	case <-ch:
		fmt.Println("done")
		return
	case <-time.After(time.Second * 5):
		fmt.Println("TimeOut")
		return
	}
}

// 1.mock normal condition
func execFuncNormal() error {
	time.Sleep(time.Millisecond * 30)
	return nil
}

// 2.mock error condition
func execFuncError() error {
	return errors.New("execFunc error")
}

// 3.mock panic condition
func execFuncPanic() error {
	panic("exceFunc panic")
}

// 3. mock TimeOut
func execFuncTimeOut() error {
	time.Sleep(time.Second * 3)
	return nil
}
