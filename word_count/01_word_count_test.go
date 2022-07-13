// fib_test.go
package main

import "testing"

func BenchmarkGetAlphanumericNumByASCII(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAlphanumericNumByASCII("108条梁山man")
	}
}

func BenchmarkGetAlphanumericNumByRegExp(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetAlphanumericNumByRegExp("108条梁山man")
	}
}
