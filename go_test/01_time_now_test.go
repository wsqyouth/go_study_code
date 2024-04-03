package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/agiledragon/gomonkey"
)

func TestPrintTime(t *testing.T) {

	// 创建一个新的 patch 集合
	patches := gomonkey.NewPatches()
	defer patches.Reset()

	// 应用 patch，mock time.Now
	patches.ApplyFunc(time.Now, func() time.Time {
		fmt.Println("Mocking time.Now")
		return time.Date(2024, time.April, 0, 0, 2, 0, 0, time.UTC)
	})

	// 现在调用 printTime 函数，它应该使用 mock 的时间
	output, err := printTime()
	if err != nil {
		t.Errorf("Not expected an error but got %v", err)
	}
	expectedOutput := "2024-03-31 00:02:00"
	if output != expectedOutput {
		t.Errorf("Expected output '%s' but got '%s'", expectedOutput, output)
	}
}
