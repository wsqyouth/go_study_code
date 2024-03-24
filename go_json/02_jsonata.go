package main

import (
	"encoding/json"
	"fmt"

	"github.com/blues/jsonata-go"
	"github.com/blues/jsonata-go/jtypes"
)

func main() {
	// 注册自定义的JSONata扩展函数
	jsonata.RegisterExts(map[string]jsonata.Extension{
		"evenNumbers": jsonata.Extension{
			Func: evenNumbers,
		}})

	// JSON数据样本
	jsonData := `{
		"name": "John Doe",
		"age": 30,
		"numbers": [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
	}`

	// JSONata表达式
	jsonataExpression := `$evenNumbers(numbers)`

	// 使用jsonata.MustCompile来编译表达式，如果表达式无效，程序将崩溃
	compiledExpr := jsonata.MustCompile(jsonataExpression)

	// 解析JSON数据
	var data interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("Error parsing JSON data:", err)
		return
	}
	fmt.Println("Parsed Data:", data)

	// 使用compiledExpr.Eval来评估数据
	res, err := compiledExpr.Eval(data)
	if err != nil {
		fmt.Println("Error evaluating expression:", err)
		return
	}

	// 将结果转换为字符串并输出
	resultArray, ok := res.([]interface{})
	if !ok {
		fmt.Println("Invalid result type")
		return
	}
	fmt.Println("Result:", resultArray)
}

// 定义一个自定义的JSONata扩展函数，用于找出数组中的偶数
func evenNumbers(data interface{}) (interface{}, error) {
	// 确保传入的是数组
	array, ok := data.([]interface{})
	if !ok {
		return jtypes.ArgUndefined(0), fmt.Errorf("expected an array")
	}

	// 过滤数组中的偶数
	var evens []interface{}
	for _, item := range array {
		// 请注意，根据 JSON 的规范，数值类型在 Go 中会被解析为 float64 类型，因此我们需要将其转换为整数进行偶数判断。
		num, ok := item.(float64)
		if !ok {
			continue
		}
		if int(num)%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens, nil
}
