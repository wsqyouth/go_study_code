package main

import (
	"fmt"
)

func main() {
	var arr []int64
	// 输入数据集
	for i := 0; i < 11; i++ {
		arr = append(arr, int64(i))
	}
	fmt.Println("src data:", arr)
	res := splitArrayByStep(arr, 5)
	for _, v := range res {
		fmt.Println(v)
	}
}

// 将数组arr按指定大小进行分隔
func splitArrayByStep(arr []int64, step int64) [][]int64 {
	max := int64(len(arr))
	var segments = make([][]int64, 0)
	quantity := max / step
	remainder := max % step
	i := int64(0)
	for i = int64(0); i < quantity; i++ {
		segments = append(segments, arr[i*step:(i+1)*step])
	}
	if remainder != 0 {
		segments = append(segments, arr[i*step:i*step+remainder])
	}
	return segments
}
