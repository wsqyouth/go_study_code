package main

import (
	"fmt"
	"math"
	"regexp"
	"unicode"
)

func main() {
	stringValue := "严苛成品样板·艺术体验中心·景观示范区，全维实景敬呈，恭迎品鉴"
	stringValueNew := "··，，"
	// minLen := 1
	// maxLen := 30
	// if IsLtLimit(stringValue, minLen) {
	// 	fmt.Printf(" < error. min_len: %v, param_len : %v, param: %v", minLen, GetMonospacedCharCount(stringValue), stringValue)
	// }
	// if IsGtLimit(stringValue, maxLen) {
	// 	fmt.Printf(" > error. maxLen: %v, param_len : %v, param: %v", maxLen, GetMonospacedCharCount(stringValue), stringValue)
	// }
	fmt.Printf("len: %v", GetStrLength(stringValue))
	fmt.Printf("len: %v", GetStrLength(stringValueNew))

}

func GetMonospacedCharDoubleCount(data string) int {
	count := 0
	for _, char := range data {
		if char >= 0 && char < 128 {
			count++
		} else {
			count += 2
		}
	}
	return count
}

// 大于
func IsGtLimit(data string, limit int) bool {
	count := GetMonospacedCharCount(data)
	fmt.Printf("len:%v, limit:%v\n", count, limit)
	return count > limit
}

// 小于
func IsLtLimit(data string, limit int) bool {
	count := GetMonospacedCharCount(data)
	fmt.Printf("len:%v, limit:%v\n", count, limit)
	return count < limit
}

func GetMonospacedCharCount(data string) int {
	doubleCount := GetMonospacedCharDoubleCount(data)
	divideValue := float64(doubleCount) / 2
	return int(math.Ceil(divideValue))
}

// GetStrLength 返回输入的字符串的字数，汉字和中文标点算 1 个字数，英文和其他字符 2 个算 1 个字数，不足 1 个算 1个
func GetStrLength(str string) int64 {
	var total float64

	reg := regexp.MustCompile("/·|，|。|《|》|‘|’|”|“|；|：|【|】|？|（|）|、|！|/")

	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) || reg.Match([]byte(string(r))) {
			total = total + 1
		} else {
			total = total + 0.5
		}
	}

	return int64(math.Ceil(total))
}
