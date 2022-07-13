package main

import "fmt"

func main() {
	fmt.Println("hello")
}

// 修改了原切片
func DeleteElemSlice(arr []int, elem int) []int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			arr = append(arr[:i], arr[i+1:]...) //打散
			i--
		}
	}
	return arr
}

// 删除指定元素
func DeleteElemSliceV2(arr []int, elem int) []int {
	tmp := make([]int, 0, len(arr))
	for _, v := range arr {
		if v != elem {
			tmp = append(tmp, v)
		}
	}
	return tmp
}
