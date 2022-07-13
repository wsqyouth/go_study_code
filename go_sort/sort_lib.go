package main

import (
	"fmt"
	"sort"
)

func main() {
	// []int 排序
	slInt := []int{5, 2, 6, 3, 1, 4} // unsorted
	sort.Ints(slInt)
	fmt.Println(slInt) // 输出 [1 2 3 4 5 6]

	// []float64 排序
	slF64 := []float64{5.2, -1.3, 0.7, -3.8, 2.6} // unsorted
	sort.Float64s(slF64)
	fmt.Println(slF64) // 输出 [-3.8 -1.3 0.7 2.6 5.2]

	// []string 字典序
	slStr := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(slStr)
	fmt.Println(slStr) // 输出 [Alpha Bravo Delta Go Gopher Grin]

	// 自定义排序
	studentSort()
}

func studentSort() {
	slStdnt := []struct {
		Name   string
		Age    int
		Height int
	}{
		{"Alice", 23, 175},
		{"David", 18, 185},
		{"Eve", 18, 165},
		{"Bob", 25, 170},
	}

	// 用 age 排序，年龄相等的元素保持原始顺序
	sort.SliceStable(slStdnt, func(i, j int) bool {
		return slStdnt[i].Age < slStdnt[j].Age
	})
	fmt.Println(slStdnt)
}

//原文链接：https://blog.csdn.net/K346K346/article/details/118314382
