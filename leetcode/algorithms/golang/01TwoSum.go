package main

import (
	"fmt"
)

func main() {
	nums :=[]int{3,2,4}
	fmt.Println(twoSum(nums,6))
}

func twoSum(nums []int, target int) []int {
	var targetIndex []int
	for i:=0; i< len(nums)-1;i++{
		for j:=i+1;j < len(nums); j++{
			if nums[i] + nums[j] == target {
				targetIndex = append(targetIndex,i,j)
				return targetIndex
			}
		}
	}
	return nil
}