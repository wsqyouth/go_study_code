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
	stu1 := Student{1, "paopao"}
	stu2 := Student{id: 2, name: "coopers"}
	stuArr := []*Student{&stu1, &stu2}
	stuInfoMap := make(map[int]*Student)
	for _, item := range stuArr {
		stuInfoMap[item.id] = item
	}
	fmt.Println(printAny(stuInfoMap))
	fmt.Println((*(stuInfoMap[2])))
	/*
		for i := 0; i < len(stuArr); i++ {
			stuInfoMap[stuArr[i].id] = stuArr[i]
		}
		fmt.Println(printAny(stuInfoMap))
	*/
}

func printAny(v interface{}) string {
	b, err := jsoniter.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", v)
	}
	return string(b)
}
