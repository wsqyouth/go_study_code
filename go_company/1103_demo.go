package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

func main() {

	makeComplexDemo()
}

type Student struct {
	id   int
	name string
}

func makeComplexDemo() {
	stuArr := []Student{Student{1, "paopao"}, Student{2, "coopers"}}
	stuInfoMap := make(map[int]Student)
	// 错误写法
	for _, item := range stuArr {
		stuInfoMap[item.id] = item
	}
	fmt.Println(printAny(stuInfoMap))
	// 正确写法
	for i := 0; i < len(stuArr); i++ {
		stuInfoMap[stuArr[i].id] = stuArr[i]
	}
	fmt.Println(printAny(stuInfoMap))
}

func printAny(v interface{}) string {
	b, err := jsoniter.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
