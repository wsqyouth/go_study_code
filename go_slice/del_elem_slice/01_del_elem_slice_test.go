// fib_test.go
package main

import "testing"

func getSlice(n int) []int {
	a := make([]int, 0, n)
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a = append(a, 0)
			continue
		}
		a = append(a, 1)
	}
	return a
}

func BenchmarkDeleteSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DeleteElemSlice(getSlice(10), 0)
	}
}
func BenchmarkDeleteSlice2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DeleteElemSliceV2(getSlice(10), 0)
	}
}
