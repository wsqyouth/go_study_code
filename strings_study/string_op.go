package main

import (
	"fmt"
	"strings"
)

func main() {
	orignSubID := "10001010;10000838"
	var dstID string
	slices := strings.SplitN(orignSubID, ";", 3)
	if len(slices) >= 2 {
		dstID = "0;" + slices[1]
	}
	fmt.Println(dstID)
}
