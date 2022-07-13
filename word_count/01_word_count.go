package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "108条梁山man"
	fmt.Println(GetAlphanumericNumByASCII(str))
	fmt.Println(GetAlphanumericNumByRegExp(str))
}

func GetAlphanumericNumByASCII(s string) int {
	num := int(0)
	for _, c := range s {
		switch {
		case '0' <= c && c <= '9':
			fallthrough
		case 'a' <= c && c <= 'z':
			fallthrough
		case 'A' <= c && c <= 'Z':
			num++
		default:
		}
	}
	return num
}

// GetAlphanumericNumByRegExp 根据正则表达式获取字母数字数量。
func GetAlphanumericNumByRegExp(s string) int {
	rNum := regexp.MustCompile(`\d`)
	rLetter := regexp.MustCompile("[a-zA-Z]")
	return len(rNum.FindAllString(s, -1)) + len(rLetter.FindAllString(s, -1))
}
