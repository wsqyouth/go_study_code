// fib_test.go
package main

import "testing"

func BenchmarkRemoveDuplicateElement(b *testing.B) {
	s := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}
	for n := 0; n < b.N; n++ {
		removeDuplicateElement(s)
	}
}
func BenchmarkRemoveDuplicateString(b *testing.B) {
	s := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}
	for n := 0; n < b.N; n++ {
		removeDuplicateString(s)
	}
}
