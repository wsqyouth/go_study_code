package main

import (
	"fmt"
	"time"
)

func main() {
	printTime()
}

func printTime() (timeStr string, err error) {
	// 获取当前时间
	nowTime := time.Now()
	// 打印当前时间
	fmt.Println(nowTime)
	return nowTime.Format("2006-01-02 15:04:05"), nil
}
