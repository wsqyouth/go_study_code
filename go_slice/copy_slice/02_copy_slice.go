package main

import (
	"fmt"
)

func main() {
	var arr []int64
	// 输入数据集
	for i := 0; i < 10; i++ {
		arr = append(arr, int64(i))
	}
	arrNew := copyArr(arr)
	fmt.Println("src data:", arrNew)
	arrNew1 := copyArr(arr)
	fmt.Println("src data:", arrNew1)
}

func copyArr(arr []int64) []int64 {
	var arrNew []int64
	for _, data := range arr {
		arrNew = append(arrNew, data)
	}
	return arrNew
}
func copyArrNew(arr []int64) []int64 {
	var arrNew []int64
	arrNew = arr[:]
	return arrNew
}
