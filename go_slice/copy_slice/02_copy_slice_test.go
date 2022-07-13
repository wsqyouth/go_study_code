// fib_test.go
package main

import "testing"

//输入数据集
func generateWithCap(n int) []int64 {
	var arr []int64
	for i := 0; i < n; i++ {
		arr = append(arr, int64(i))
	}
	return arr
}

func BenchmarkCopyArr(b *testing.B) {
	b.StopTimer()
	nums := generateWithCap(10000)
	b.StartTimer() //不计入耗时
	for n := 0; n < b.N; n++ {
		copyArr(nums)
	}
}
func BenchmarkCopyArrNew(b *testing.B) {
	b.StopTimer()
	nums := generateWithCap(10000)
	b.StartTimer() //不计入耗时
	for n := 0; n < b.N; n++ {
		copyArrNew(nums)
	}
}
