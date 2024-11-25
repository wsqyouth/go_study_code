package main

import (
	"fmt"
	"time"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	start := time.Now() // 记录开始时间
	err := validatePdf()
	elapsed := time.Since(start)                // 计算耗时
	fmt.Printf("Execution time: %s\n", elapsed) // 输出执行耗时

	if err != nil {
		fmt.Print(err)
	}
}
func validatePdf() error {
	// filePath := "test_golang.pdf"
	filePath := "test_normal.pdf"
	// filePath := "ab82c5da-487b-4a98-bb53-e0d6dd9b7649-1725343467291.pdf"

	// 使用 pdfcpu 库校验 PDF 文件
	err := api.ValidateFile(filePath, nil)
	if err != nil {
		// fmt.Println("The PDF is invalid and cannot be opened.")
		return err
	}
	// fmt.Println("The PDF is valid and can be opened.")
	return nil
}
