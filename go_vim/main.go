package main

import "fmt"

func main() {

	fmt.Printf("vim-go:%v",Add(30,20))
}

func Add(a, b int64) int64 {
	return a+b
}
