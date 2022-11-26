package main

import (
	"fmt"
)

func main() {
	studentArr := []Student{Student{1, "pony"}, Student{2, "mark"}}
	newStudentArr := []Student{Student{1, "paopao"}, Student{3, "davis"}}

	// 建立映射
	studentMap := map[int]int{}
	for index, val := range studentArr {
		studentMap[val.id] = index
	}

	// 存在则合并,不存在则更新
	for _, newVal := range newStudentArr {
		if index, ok := studentMap[newVal.id]; ok {
			studentArr[index] = newVal
		} else {
			studentArr = append(studentArr, newVal)
		}
	}
	fmt.Println(studentArr)
}

type Student struct {
	id   int
	name string
}
