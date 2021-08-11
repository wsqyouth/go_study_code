package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var sum uint64
	idSet := []uint64{21}
	for _, v := range idSet {
		sum += uint64(math.Pow(2, float64(v)-1))
	}
	fmt.Println("sum", sum)
	fmt.Printf("%b\n", sum)

	b := reverse(strconv.FormatInt(int64(sum), 2))
	fmt.Println(b)

	var outIdSet []uint64
	for k, v := range b {
		if v == '1' {
			outIdSet = append(outIdSet, uint64(k+1))
		}
	}
	fmt.Println(outIdSet)
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}
