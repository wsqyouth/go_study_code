package main

import (
	"fmt"
)

func main() {
	s := []string{"hello", "world", "hello", "golang", "hello", "ruby", "php", "java"}

	fmt.Println(removeDuplicateElement(s))
	fmt.Println(removeDuplicateString(s))
}

func removeDuplicateElement(arr []string) []string {
	res := make([]string, 0, len(arr))
	tempMap := map[string]struct{}{}
	for _, item := range arr {
		if _, ok := tempMap[item]; !ok {
			tempMap[item] = struct{}{}
			res = append(res, item)
		}
	}
	return res
}

func removeDuplicateString(strList []string) []string {
	if len(strList) == 0 {
		return nil
	}
	var strMap = make(map[string]bool, len(strList))
	duplicateStrList := []string{}
	for _, str := range strList {
		if _, ok := strMap[str]; ok {
			continue
		}
		strMap[str] = true
		duplicateStrList = append(duplicateStrList, str)
	}
	return duplicateStrList
}

//使用map去重，空struct不占用空间
