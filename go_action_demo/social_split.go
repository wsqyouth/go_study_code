package main

import "fmt"

func main() {
	templateSoicalMap := make(map[int32][]int32)
	//socialSkill := []int32{311, 480, 482, 516, 519, 539, 540, 599, 618, 641, 642, 643, 698, 699, 721, 1064, 1065, 1465, 1480, 1707, 1708, 1733, 1748, 1765, 1766, 1814}
	firstCommet := []int32{311, 480, 482, 516, 519, 539, 540, 599, 618, 641, 642, 643, 698, 699, 721, 1064, 1065, 1465, 1480, 1707, 1708, 1733, 1748, 1765, 1766, 1814}
	dataMonic := []int32{482, 519, 540, 599, 618, 698, 721, 1064, 1065, 1465, 1480, 1708, 1733, 1748, 1765, 1814}
	gestureSpec := []int32{135, 311, 480, 482, 484, 516, 519, 538, 539, 540, 555, 556, 559, 560, 599, 618, 641, 642, 643, 698, 699, 701, 708, 721, 913, 926, 928, 972, 998, 1003, 1064, 1065, 1465, 1480, 1515, 1707, 1708, 1733, 1765, 1766, 1814, 1817, 1861}
	//服务号对话能力 0 0 1
	initSlice := []int32{0, 0, 1}
	for _, data := range gestureSpec {
		val, ok := templateSoicalMap[data]
		if ok {
			fmt.Println("error", val)
		} else {
			templateSoicalMap[data] = initSlice
		}
	}

	fmt.Println(templateSoicalMap)
	//dataMonic 0 1 0
	for _, data := range dataMonic {
		_, ok := templateSoicalMap[data]
		if ok {
			templateSoicalMap[data] = []int32{0, 1, 1}
		}
	}

	//firstCommet  1 0 0
	for _, data := range firstCommet {
		_, ok := templateSoicalMap[data]
		if ok {
			templateSoicalMap[data] = []int32{1, 1, 1}
		}
	}

	//for social := range templateSoicalMap {
	//	fmt.Println(social, templateSoicalMap[social][0], templateSoicalMap[social])
	//}

	//templadArr := []int32{1707, 1708, 599, 698, 699, 480, 482, 1748, 925, 926, 927, 928, 972, 484, 608, 556} //coopers
	//templadArr := []int32{1065, 1733, 1814, 555, 559, 560, 701, 910, 1515, 929} //coopers
	//templadArr := []int32{1765, 1766}                                               //coopers
	//templadArr := []int32{567, 998, 708, 913, 1817, 1818, 1861, 516, 519, 539, 540} //lynlli
	templadArr := []int32{135, 538, 1003, 1064, 1465, 1480, 721, 618, 641, 642, 643, 311} //borough
	for _, data := range templadArr {
		val, ok := templateSoicalMap[data]
		if ok {
			fmt.Println("found val", data, val)
		} else {
			fmt.Println("not found", data)
		}
	}
	//fmt.Println(templateSoicalMap)
}
