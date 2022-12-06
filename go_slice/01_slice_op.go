package main

import "fmt"

func main() {
	res := make([]int, 0)
	// 切片尾部追加元素
	for i := 0; i < 10; i++ {
		res = append(res, i)
	}
	// 切片特定位置插入元素
	rear := append([]int{}, res[3:]...)
	res = append(res[0:3], 100)
	res = append(res, rear...)
	fmt.Println(res)
	// 切片copy操作
}
