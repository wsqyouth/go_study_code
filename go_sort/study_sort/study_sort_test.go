// fib_test.go
package main

import (
	"math/rand"
	"testing"
	"time"
)

//输入数据集
func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func BenchmarkBubbleSort(b *testing.B) {
	b.StopTimer()
	nums := generateWithCap(10000)
	b.StartTimer() //不计入耗时
	for n := 0; n < b.N; n++ {
		BubbleSort(nums)
	}
}
func BenchmarkSelectSort(b *testing.B) {
	b.StopTimer()
	nums := generateWithCap(10000)
	b.StartTimer() //不计入耗时
	for n := 0; n < b.N; n++ {
		SelectSort(nums)
	}
}
func BenchmarkQuickSort(b *testing.B) {
	b.StopTimer()
	nums := generateWithCap(10000)
	b.StartTimer() //不计入耗时
	for n := 0; n < b.N; n++ {
		QuickSort(nums, 0, len(nums)-1)
	}
}
