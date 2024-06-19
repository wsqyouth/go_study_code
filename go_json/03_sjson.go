package main

import (
	"fmt"

	"github.com/tidwall/sjson"
)

func main() {
	jsonStr := `
	{
		"name": {"first": "Tom", "last": "Anderson"},
		"age":37,
		"children": ["Sara","Alex","Jack"],
		"fav.movie": "Deer Hunter",
		"friends": [
		  {"first": "James", "last": "Murphy"},
		  {"first": "Roger", "last": "Craig"}
		],
		"address": {   
		}
	  }
	`
	value, _ := sjson.Set(jsonStr, "name.last", "Anderson")
	println(value)

	addressStr := `
	{
	    "street": "1234567890",
	    "city": "New York",
	    "state": "NY"
	}`

	// 设置 json 对象
	value, _ = sjson.Set(jsonStr, "address", addressStr)
	println(value)
	// 设置 json 对象时 需要使用这种方式
	value, _ = sjson.SetRaw(jsonStr, "address", addressStr)
	fmt.Println(value)
}

// 文章参考: https://juejin.cn/post/7150823363055714334
/*
思考: 使用 sjson, gjson, jsonata 组合能够实现非常强大的功能
*/
