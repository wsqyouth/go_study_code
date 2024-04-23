package main

import (
	"fmt"
	"time"
)

func main() {
	tz := "Australia/Sydney"
	location, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	date := "2024-04-02" // 定义日期
	t := "08:47:00"      // 定义时间

	// 将日期和时间拼接为一个完整的字符串
	parsedTime, err := time.ParseInLocation("2006-01-02 15:04:05", date+" "+t, location)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	// 将时区转换为 UTC
	utcTime := parsedTime.UTC()

	// 输出UTC时间
	fmt.Printf("UTC Time: %s\n", utcTime.Format("2006-01-02 15:04:05"))

	// 输出UTC时间的YYYYMMDD格式
	fmt.Printf("UTC Time (YYYYMMDD format): %s\n", utcTime.Format("20060102"))
}
