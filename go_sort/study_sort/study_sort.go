package main

import "fmt"

func main() {
	numbers := []int{6, 2, 7, 5, 8, 9}
	SelectSort(numbers)
	fmt.Println(numbers)
	numbers = []int{6, 2, 7, 5, 8, 9}
	BubbleSort(numbers)
	fmt.Println(numbers)
	numbers = []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
	QuickSort(numbers, 0, len(numbers)-1)
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

func partition(list []int, low, high int) int {
	pivot := list[low] //导致 low 位置值为空
	for low < high {
		//high指针值 >= pivot high指针👈移
		for low < high && pivot <= list[high] {
			high--
		}
		//填补low位置空值
		//high指针值 < pivot high值 移到low位置
		//high 位置值空
		list[low] = list[high]
		//low指针值 <= pivot low指针👉移
		for low < high && pivot >= list[low] {
			low++
		}
		//填补high位置空值
		//low指针值 > pivot low值 移到high位置
		//low位置值空
		list[high] = list[low]
	}
	//pivot 填补 low位置的空值
	list[low] = pivot
	return low
}

func QuickSort(list []int, low, high int) {
	if high > low {
		//位置划分
		pivot := partition(list, low, high)
		//左边部分排序
		QuickSort(list, low, pivot-1)
		//右边排序
		QuickSort(list, pivot+1, high)
	}
}
