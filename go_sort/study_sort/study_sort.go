package main

import "fmt"

func main() {
	/*
		numbers := []int{6, 2, 7, 5, 8, 9}
		SelectSort(numbers)
		fmt.Println(numbers)
		numbers = []int{6, 2, 7, 5, 8, 9}
		BubbleSort(numbers)
		fmt.Println(numbers)
		numbers = []int{2, 44, 4, 8, 33, 1, 22, -11, 6, 34, 55, 54, 9}
		QuickSort(numbers, 0, len(numbers)-1)
		fmt.Println(numbers)
	*/
	HeapSort()
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

// 本例为最小堆
// 最大堆只需要修改less函数即可
type Heap []int

func (h Heap) swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) less(i, j int) bool {
	return h[i] < h[j]
}
func (h Heap) up(i int) {
	for {
		f := (i - 1) / 2 // 父亲结点
		if i == f || h.less(f, i) {
			break
		}
		h.swap(f, i)
		i = f
	}
}

// 注意go中所有参数转递都是值传递
// 所以要让h的变化在函数外也起作用，此处得传指针
func (h *Heap) Push(x int) {
	*h = append(*h, x)
	h.up(len(*h) - 1)
}

func (h Heap) down(i int) {
	for {
		l := 2*i + 1 // 左孩子
		if l >= len(h) {
			break // i已经是叶子结点了
		}
		j := l
		if r := l + 1; r < len(h) && h.less(r, l) {
			j = r // 右孩子
		}
		if h.less(i, j) {
			break // 如果父结点比孩子结点小，则不交换
		}
		h.swap(i, j) // 交换父结点和子结点
		i = j        //继续向下比较
	}
}

// 删除堆中位置为i的元素
// 返回被删元素的值
func (h *Heap) Remove(i int) (int, bool) {
	if i < 0 || i > len(*h)-1 {
		return 0, false
	}
	n := len(*h) - 1
	h.swap(i, n) // 用最后的元素值替换被删除元素
	// 删除最后的元素
	x := (*h)[n]
	*h = (*h)[0:n]
	// 如果当前元素大于父结点，向下筛选
	if (*h)[i] > (*h)[(i-1)/2] {
		h.down(i)
	} else { // 当前元素小于父结点，向上筛选
		h.up(i)
	}
	return x, true
}

// 弹出堆顶的元素，并返回其值
func (h *Heap) Pop() int {
	n := len(*h) - 1
	h.swap(0, n)
	x := (*h)[n]
	*h = (*h)[0:n]
	h.down(0)
	return x
}

func (h Heap) Init() {
	n := len(h)
	// i > n/2-1 的结点为叶子结点本身已经是堆了
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i)
	}
}

// 学习堆排序的相关操作
// ref: https://www.cnblogs.com/yahuian/p/go-heap.html
func HeapSort() {
	h := Heap{20, 7, 3, 10, 15, 25, 30, 17, 19}
	h.Init()
	fmt.Println(h) // [3 7 20 10 15 25 30 17 19]

	h.Push(6)
	fmt.Println(h) // [3 6 20 10 7 25 30 17 19 15]

	x, ok := h.Remove(5)
	fmt.Println(x, ok, h) // 25 true [3 6 15 10 7 20 30 17 19]

	y, ok := h.Remove(1)
	fmt.Println(y, ok, h) // 6 true [3 7 15 10 19 20 30 17]

	z := h.Pop()
	fmt.Println(z, h) // 3 [7 10 15 17 19 20 30]
}
