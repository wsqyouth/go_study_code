package main

import (
	"fmt"

	"github.com/tidwall/gjson"
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
	value := gjson.Get(jsonStr, "name.last")
	fmt.Println(value)

	value = gjson.Get(jsonStr, "friends.#.first")
	fmt.Println(value)

	// 递归查询
	nameList := []string{}
	gjson.Get(jsonStr, "friends.#.first").ForEach(func(key, value gjson.Result) bool {
		nameList = append(nameList, value.String())
		return true
	})
	fmt.Println(nameList)
}

// 参考: https://juejin.cn/post/7150651352057249822/
