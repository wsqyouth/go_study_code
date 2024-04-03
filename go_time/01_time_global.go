package main

import (
	"fmt"
	"time"
)

// 映射表：ISO3 代码与 IANA 时区名称的对应关系
var countryTimezones = map[string]string{
	"USA": "America/New_York",
	// 添加其他国家的 ISO3 代码和对应的 IANA 时区名称
}

func main() {
	iso3CountryCode := "USA" // 替换为你想要查询的国家的 ISO3 代码

	// 获取国家对应的 IANA 时区名称
	ianaTimezone, err := getTimezone(iso3CountryCode)
	if err != nil {
		fmt.Printf("无法找到国家的时区：%s\n", err)
		return
	}

	// 获取当前时间的本地时间
	localTime := time.Now().In(ianaTimezone)

	// 格式化为字符串类型
	localTimeString := localTime.Format("2006-01-02 15:04:05")

	// 打印本地时间
	fmt.Println(localTimeString)
}

// 获取国家对应的时区信息
func getTimezone(iso3CountryCode string) (*time.Location, error) {
	ianaTimezone, ok := countryTimezones[iso3CountryCode]
	if !ok {
		return nil, fmt.Errorf("无法找到国家的时区")
	}

	return time.LoadLocation(ianaTimezone)
}
