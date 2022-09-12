package main

import "testing"

func BenchmarkCheckParamSample(b *testing.B) {
	var req ReqParam
	req.id = 23
	//req.name = "hello"
	for n := 0; n < b.N; n++ {
		checkParamSample(req)
	}
}
func BenchmarkCheckParam(b *testing.B) {
	var req ReqParam
	req.id = 23
	//req.name = "hello"
	for n := 0; n < b.N; n++ {
		checkParam(req)
	}
}
