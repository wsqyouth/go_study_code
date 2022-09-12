package main

import "fmt"

type Foo interface {
	foo()
}

type A struct {
}

func (a A) foo() {
	fmt.Println("A foo")
}

func (a A) bar() {
	fmt.Println("A bar")
}

func callFoo(f Foo) {
	f.foo()
}

/*
func main() {
	var a A
	callFoo(a)
}
*/
