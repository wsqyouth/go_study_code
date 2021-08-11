package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	strSplit()
	strToFloat()
}

func strSplit() {
	orignSubID := "10001010;10000838"
	var dstID string
	slices := strings.SplitN(orignSubID, ";", 3)
	if len(slices) >= 2 {
		dstID = "0;" + slices[1]
	}
	fmt.Println(dstID)
}
func strToFloat() {
	str := "12.345"
	num, err := strconv.ParseFloat(str, 10)
	if err != nil {
		fmt.Printf("strtofloat err. err: %v", err)
	}
	fmt.Printf("%T  %v\n", num, num)
}
