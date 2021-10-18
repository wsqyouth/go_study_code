package main

import "fmt"

func main() {
	numbers := []int{6, 2, 7, 5, 8, 9}
	SelectSort(numbers)
	fmt.Println(numbers)
	numbers = []int{6, 2, 7, 5, 8, 9}
	BubbleSort(numbers)
	fmt.Println(numbers)
}

func BubbleSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		for j := i + 1; j < length; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
}

func SelectSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length; i++ {
		min := i //index
		for j := length - 1; j > i; j-- {
			if arr[j] < arr[i] {
				min = j
			}
		}

		//swap
		arr[i], arr[min] = arr[min], arr[i]
	}
}
