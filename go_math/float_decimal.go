package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	price, err := decimal.NewFromString("136.02")
	if err != nil {
		panic(err)
	}

	quantity := decimal.New(3, 0) // value * 10 ^ exp = 3 * 10 ^ 0 = 3

	fee, _ := decimal.NewFromString(".035")       // 0.035
	taxRate, _ := decimal.NewFromString(".08875") // 0.08875

	subtotal := price.Mul(quantity) // 136.02 * 3 = 408.06

	preTax := subtotal.Mul(fee.Add(decimal.NewFromFloat(1))) // 408.06 * 1.035 = 422.3421

	total := preTax.Mul(taxRate.Add(decimal.NewFromFloat(1))) // 422.3421 * 1.08875 = 459.824961375

	fmt.Println("Subtotal:", subtotal)                      // Subtotal: 408.06
	fmt.Println("Pre-tax:", preTax)                         // Pre-tax: 422.3421
	fmt.Println("Taxes:", total.Sub(preTax))                // Taxes: 37.482861375
	fmt.Println("Total:", total)                            // Total: 459.824961375
	fmt.Println("Tax rate:", total.Sub(preTax).Div(preTax)) // Tax rate: 0.08875
}

/*
学习下 github 中浮点数的库函数使用
*/
