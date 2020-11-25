package main

import "fmt"

func main() {
	var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	fmt.Println("type:%T,value:%V", p, p)
	fmt.Println("type:%T,value:%V", v, v)
	// Unnecessarily complex:
	var pp *[]int = new([]int)
	*pp = make([]int, 100, 100)

	// Idiomatic:
	vv := make([]int, 100)

	fmt.Println("vv:%T", vv)
}

/*
make 函数只能用来创建切片、映射（map） 和 信道，返回的是一个被初始化、类型为 T 的值（内存不是全零）。造成这种区别的原因是，切片、映射和信道 这三个类型的数据底层引用了其他的数据结构，而这些底层的数据结构在使用前必须先初始化才可以工作。



make([]int,10,100)
初始化了一个包含 100 个整数的数组，同时创建了一个长度为 10、容量为 100 、指向这个数组 的前 10 个元素的切片

new([]int)
返回了一个新创建且被置零的切片的指针，指针指向的是值为 nil 的切片值

*/
