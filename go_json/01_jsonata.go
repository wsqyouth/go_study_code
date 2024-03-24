package main

import (
	"encoding/json"
	"fmt"
	"log"

	jsonata "github.com/blues/jsonata-go"
)

const jsonString = `
    {
        "orders": [
            {"price": 10, "quantity": 3},
            {"price": 0.5, "quantity": 10},
            {"price": 100, "quantity": 1}
        ],
		"name": "John Doe",
		"age": 30,
		"verified": true,
		"addresses": [
			{"city": "San Francisco", "state": "CA"},
			{"city": "New York", "state": "NY"}
		]
    }
`

func main() {

	var data interface{}

	// Decode JSON.
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		log.Fatal(err)
	}

	// Create expression.
	// e := jsonata.MustCompile("orders[0].(price*quantity)")
	jsonataExpression := `addresses[0].city`
	e := jsonata.MustCompile(jsonataExpression)

	// 使用compiledExpr.Eval来评估数据
	res, err := e.Eval(data)
	if err != nil {
		log.Fatal(err)
	}

	// 输出结果
	fmt.Println(res)
}

// 参考：https://github.com/blues/jsonata-go
// https://pkg.go.dev/github.com/xiatechs/jsonata-go#section-readme
