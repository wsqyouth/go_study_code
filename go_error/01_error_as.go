package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	//checkErrorIs()
	//checkErrorAs()
	checkErrorAsOnWrap()
}

//As panics if target is not a non-nil pointer to either a type that implements error, or to any interface type.

func checkErrorAs() {
	if _, err := os.Open("non_existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}
func checkErrorIs() {
	if _, err := os.Open("non_existing"); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("Failed does not exist")
		} else {
			fmt.Println(err)
		}
	}
}

func checkErrorAsOnWrap() {
	var err = &MyError{"my error type"}
	err1 := fmt.Errorf("wrap err1: %w", err)
	err2 := fmt.Errorf("wrap err2: %w", err1)
	var e *MyError
	if errors.As(err2, &e) {
		println("MyError is on the chain of err2 ")
		println(e == err)
	} else {
		println("MyError is not on the chain of err2 ")
	}

}

type MyError struct {
	e string
}

func (e *MyError) Error() string {
	return e.e
}

//ref: https://tonybai.com/2019/10/18/errors-handling-in-go-1-13/
//思考：这里回头再来看老白的书,才获得其奥妙
//自go1.13版本后的Is和As方法会递归判断处理,这种场景可以处理被wrap的场景
