package main

import "fmt"

func main() {
	decorator(hello)("hello coopers")
}

func decorator(f func(s string)) func(s string) {
	return func(s string) {
		fmt.Println("start")
		f(s)
		fmt.Println("end")
	}
}
func hello(s string) {
	fmt.Println(s)
}
