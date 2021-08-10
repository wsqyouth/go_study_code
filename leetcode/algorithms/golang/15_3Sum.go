package main

import (
	"fmt"
)

//Input: nums =
//Output: [[-1,-1,2],[-1,0,1]]
func main() {
	nums :=[]int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) [][]int {
	var target [][]int

	for i:=0; i< len(nums)-2;i++{
		for j:=i+1;j < len(nums)-1; j++{
			for k:=j+1; k <len(nums);k++{
				if nums[i] + nums[j] + nums[k] == 0 {
					var targetTemp []int
					var isExist bool
					targetTemp= append(targetTemp,nums[i],nums[j],nums[k])
					//fmt.Println(targetTemp)
					for _, item :=range target{
						fmt.Printf("%v--%v\n",item,targetTemp)
						//if SliceEqual(item,targetTemp){
						//	isExist = true
						//	break
						//}
					}
					if !isExist{
						target = append(target,targetTemp)
					}
				}
			}
		}
	}
	return target
}

func SliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

