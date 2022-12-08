package main

import "fmt"

func main() {
	numStrSlice := []string{"1", "2", "3", "4", "5", "6"}
	subNumStrSlice := []string{"1", "3", "5"}
	// 查找
	val, ok := Find(subNumStrSlice, "3")
	if ok {
		fmt.Println("found", val)
	} else {
		fmt.Println("not found")
	}
	// 交集
	intersectArray := IntersectArray(numStrSlice, subNumStrSlice)
	fmt.Println("intersectArray:", intersectArray)
	// 差集
	diffArray := DiffArray(numStrSlice, subNumStrSlice)
	fmt.Println(diffArray)
}

// Find 获取一个切片并在其中查找元素。如果找到它，它将返回它的val，否则它将返回-1和一个错误的bool。
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// IntersectArray 求两个切片的交集
func IntersectArray(a []string, b []string) []string {
	var inter []string
	mp := make(map[string]bool)

	for _, s := range a {
		if _, ok := mp[s]; !ok {
			mp[s] = true
		}
	}
	for _, s := range b {
		if _, ok := mp[s]; ok {
			inter = append(inter, s)
		}
	}

	return inter
}

// DiffArray 求两个切片的差集
func DiffArray(a []string, b []string) []string {
	var diffArray []string
	temp := map[string]struct{}{}

	for _, val := range b {
		if _, ok := temp[val]; !ok {
			temp[val] = struct{}{}
		}
	}

	for _, val := range a {
		if _, ok := temp[val]; !ok {
			diffArray = append(diffArray, val)
		}
	}

	return diffArray
}

// RemoveRepeatedElement 切片去重实现
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

//切片去重实现
func arrayUnique(arr []string) []string {
	result := make([]string, 0, len(arr))
	temp := map[string]struct{}{}
	for i := 0; i < len(arr); i++ {
		if _, ok := temp[arr[i]]; ok != true {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}
