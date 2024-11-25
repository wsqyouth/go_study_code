package main

import (
	"testing"
)

func BenchmarkValidatePdf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := validatePdf()
		if err != nil {
			b.Error(err) // 如果有错误，记录错误
		}
	}
}
