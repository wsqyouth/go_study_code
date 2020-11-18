package main

import "fmt"

func showCase(str string) {
	switch str {
	case "A":
		fmt.Println("match A")
	case "B":
		fmt.Println("match B")
	case "C":
		fmt.Println("match C")
	default:
		fmt.Println("other")
	}
}

func showCaseNoSensive(str string) {
	switch str {
	case "A", "a":
		fmt.Println("match A/a")
	case "B", "b":
		fmt.Println("match B/b")
	case "C", "c":
		fmt.Println("match C/c")
	default:
		fmt.Println("other")
	}
}

func main() {
	fmt.Println("Hello, World!")
	showCase("A")
	showCase("a")

	//测试兼容情
	showCaseNoSensive("A")
	showCaseNoSensive("a")
}
