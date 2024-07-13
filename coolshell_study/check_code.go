package main

import "fmt"

/*
 * 检查码（Modules 10/ Weight 3）之计算方法
 *详见：https://drive.google.com/file/d/1lqwnmWEqfnO7mksRnAkS_S75KR0fQ2bW/view
 */
func calculateCheckCode(trackingNumberPrefix string) int {
	// 1. 拆分 tracking_number_prefix 为单个数字数组
	digits := make([]int, len(trackingNumberPrefix))
	for i, char := range trackingNumberPrefix {
		digits[i] = int(char - '0')
	}

	// 2. 分成奇数位数组和偶数位数组
	var oddSum, evenSum int
	for i := len(digits) - 1; i >= 0; i-- {
		if (len(digits)-i)%2 == 1 { // 奇数位
			oddSum += digits[i]
		} else { // 偶数位
			evenSum += digits[i]
		}
	}

	// 3. 计算权重
	oddSum *= 3
	totalSum := oddSum + evenSum

	// 4. 计算检查码
	checkCode := (10 - totalSum%10) % 10

	return checkCode
}

func main() {
	trackingNumberPrefix := "1256791022905822100"
	expectedResult := 9

	result := calculateCheckCode(trackingNumberPrefix)

	fmt.Printf("Expected Result: %d\n", expectedResult)
	fmt.Printf("Calculated Result: %d\n", result)
}
