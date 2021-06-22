package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	nowTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	sixtyDayTime := time.Now().Add(60 * 24 * time.Hour)
	fmt.Println("today date: ", nowTime)
	fmt.Println("60s days later: ", sixtyDayTime)
}
