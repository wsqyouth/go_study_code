package main

import (
	"fmt"
	"time"
)

func main() {
	//errDemo()
	//normalDemo()
	normalDemo1()
	time.Sleep(time.Second)
}

func errDemo() {
	for i := 0; i < 10; i++ {
		go func() {
			// Here i is a "free" variable, since it wasn't declared
			// as an explicit parameter of the func literal,
			// so IT'S NOT copied by value as one may infer. Instead,
			// the "current" value of i
			// (in most cases the last value of the loop) is used
			// in all the go routines once they are executed.
			processValue(i)
		}()
	}
}

func normalDemo() {

	for i := 0; i < 10; i++ {
		go func(differentI int) {
			processValue(differentI)
		}(i) // Here i is effectively passed by value since it was
		// declared as an explicit parameter of the func literal
		// and is taken as a different "differentI" for each
		// go routine, no matter when the go routine is executed
		// and independently of the current value of i.
	}
}

func normalDemo1() {
	for i := 0; i < 10; i++ {
		differentI := i //golang captured by func literal
		go func() {
			processValue(differentI)
		}()
	}
}

func processValue(i int) {
	fmt.Println(i)
}
