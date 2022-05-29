package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

func main() {
	checkErrorIs()
	checkErrorAs()
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

//ref: https://tonybai.com/2019/10/18/errors-handling-in-go-1-13/
