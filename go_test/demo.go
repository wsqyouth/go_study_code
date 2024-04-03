package main

import (
	"context"
	"fmt"
	"os"
	"strings"

	"go_test/datasource"
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

// GetDataFromDB  测试包调用函数
func GetDataFromDB(ctx context.Context) (string, error) {
	srcData, err := datasource.GetData(ctx)
	fmt.Println("get data from db", srcData, err)
	return srcData, err
}

// GetDataFromDBStruct  测试包调用方法
func GetDataFromDBStruct(ctx context.Context) (string, error) {
	ds := datasource.NewDataSource(ctx)
	srcData, err := ds.GetData(ctx)
	fmt.Println("get data from db", srcData, err)
	return srcData, err
}
