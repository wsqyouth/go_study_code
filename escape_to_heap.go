package main

import "fmt"

/*
type Person struct {
	Name string
	Age  int
}

func PersonRegister(name string, age int) *Person {
	// 场景1: 指针逃逸 位置: new(Person) escapes to heap
	p := new(Person) //局部变量s逃逸到堆

	p.Name = name
	p.Age = age

	return p
}

func main() {
	person := PersonRegister("微客鸟窝", 18)
	fmt.Println(person) //场景2: 该变量作为实参传递给 fmt.Println()，但是因为 fmt.Println() 的参数类型定义为 interface{}，因此也发生了逃逸。
}

*/
/*
func generate8191() {
	nums := make([]int, 8191) // < 64KB  make([]int, 8191) does not escape
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

func generate8192() {
	nums := make([]int, 8192) // = 64KB make([]int, 8192) does not escape
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	// 场景3: 当切片占用内存超过一定大小，或无法确定当前切片长度时，对象占用内存将在堆上分配。
	nums := make([]int, n) // 不确定大小 make([]int, n) escapes to heapu
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

func main() {
	generate8191()
	generate8192()
	generate(1)
}
*/

func Increase() func() int {
	n := 0
	// 场景4: 闭包让你可以在一个内层函数中访问到其外层函数的作用域。该闭包函数访问了外部变量 n，那变量 n 将会一直存在，直到 in 被销毁
	return func() int { // func literal escapes to heap
		n++
		return n
	}
}

func main() {
	in := Increase()
	fmt.Println(in()) // 1
	fmt.Println(in()) // 2
}

/*
go build -gcflags=-m escape_to_heap.go
参考: https://geektutu.com/post/hpg-escape-analysis.html
*/
