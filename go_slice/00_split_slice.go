package main

import (
	"fmt"
)

func main() {
	var arr []int64
	// 输入数据集
	for i := 0; i < 1011; i++ {
		arr = append(arr, int64(i))
	}
	fmt.Println("src data:", arr)
	res := splitArrayByStep(arr, 5)
	for _, v := range res {
		fmt.Println(v)
	}

	fmt.Println("-------------")
	processArrayByStep(arr,5)
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

// 将数组按指定大小的size进行分批处理
func processArrayByStep(arr []int64, batchSize int) {

	arrLen := len(arr)
	if arrLen <= batchSize{
		fmt.Printf("one process:%v\n", arr)
	}else{
		arrList := arr
		for i:=0; i*batchSize < arrLen; i++{
			startIndex := i*batchSize
			endIndex := startIndex + batchSize
			if endIndex >= arrLen {
				arr = arrList[startIndex:]
			}else{
				arr = arrList[startIndex:endIndex]
			}
			fmt.Printf("batch process:%v\n", arr)
		}
	}
}