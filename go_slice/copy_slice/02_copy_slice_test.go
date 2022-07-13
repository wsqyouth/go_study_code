// fib_test.go
package main

import "testing"

func BenchmarkCopyArr(b *testing.B) {
	var arr []int64 //输入数据集
	for i := 0; i < 10; i++ {
		arr = append(arr, int64(i))
	}
	for n := 0; n < b.N; n++ {
		copyArr(arr)
	}
}
func BenchmarkCopyArrNew(b *testing.B) {
	var arr []int64 //输入数据集
	for i := 0; i < 10; i++ {
		arr = append(arr, int64(i))
	}
	for n := 0; n < b.N; n++ {
		copyArrNew(arr)
	}
}
