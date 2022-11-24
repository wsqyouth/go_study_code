package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("hello")
}

// IsInUint32 uint32的isin
func IsInUint32(a uint32, array []uint32) bool {
	for _, b := range array {
		if a == b {
			return true
		}
	}
	return false
}

func Init() {

	key := "key_demo"
	val, ok := os.LookupEnv(key)
	if !ok {
		fmt.Println("not found")
	}
	fmt.Println("found", val)
}

type Stack struct {
	length int
}

func NewStack() *Stack {
	return &Stack{0}
}

func (s *Stack) Len() int {
	return s.length
}

// 一个全局变量，实际情况中会自动读取配置文件并初始化，但在单测中不会如此
var Map map[string]bool

// 待测函数
func IsTag(originMonomial string) bool {
	monomial := strings.TrimSuffix(originMonomial, "_7")
	if _, ok := Map[monomial]; ok {
		return false
	}
	return true
}
