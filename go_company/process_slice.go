package main

import (
	"fmt"
	"strconv"
)

func IntToSringArr(src []int)(dst []string){
	for _,ele := range src{
		str := strconv.Itoa(ele)
		dst = append(dst,str)
	}
	return dst
}
func main() {
	//srcMp :=[]int{311,460,480,482,599,618,641,642,643,698,699,721,1064,1065,1465,1480,1707,1708,1733,1748,1765,1766,1814}
	//srcGdt :=[]int{311,618,641,642,643,720,721,1064,1465,1480,1530,1531,1707,1708}

	srcMp :=[]int{1,2,3,4,5,6}
	srcGdt :=[]int{2,3,7,15}
	dstMp := IntToSringArr(srcMp)
	dstGdt := IntToSringArr(srcGdt)

	fmt.Printf("slice1长度为：%v, slice2长度为：%v\n",len(dstMp),len(dstGdt))

	in := intersect(dstMp, dstGdt)
	fmt.Println("slice1与slice2的交集为：", len(in))
	fmt.Println(in)

	di := difference(dstMp, dstGdt)
	fmt.Println("slice1与slice2的差集为：", len(di))
	fmt.Println(di)
}


//求并集
func union(slice1, slice2 []string) []string {
	m := make(map[string]int)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 0 {
			slice1 = append(slice1, v)
		}
	}
	return slice1
}

//求交集
func intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}