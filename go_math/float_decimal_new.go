package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

// Add 精确加法运算
func Add(a, b float64) float64 {
	ad := decimal.NewFromFloat(a)
	bd := decimal.NewFromFloat(b)
	result, _ := ad.Add(bd).Float64()
	return result
}

// Subtract 精确减法运算
func Subtract(a, b float64) float64 {
	ad := decimal.NewFromFloat(a)
	bd := decimal.NewFromFloat(b)
	result, _ := ad.Sub(bd).Float64()
	return result
}

// Multiply 精确乘法运算
func Multiply(a, b float64) float64 {
	ad := decimal.NewFromFloat(a)
	bd := decimal.NewFromFloat(b)
	result, _ := ad.Mul(bd).Float64()
	return result
}

// Divide 精确除法运算
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	ad := decimal.NewFromFloat(a)
	bd := decimal.NewFromFloat(b)
	result, _ := ad.Div(bd).Float64()
	return result, nil
}

// Round 精确四舍五入实现
func Round(num float64, places int) float64 {
	d := decimal.NewFromFloat(num)
	result, _ := d.Round(int32(places)).Float64()
	return result
}

// FormatFloat 保持原样
func FormatFloat(num float64, places int) string {
	if places < 0 {
		places = 0
	}
	return strconv.FormatFloat(num, 'f', places, 64)
}

// main函数保持不变
func main() {
	a := Add(1, 0.2)
	fmt.Printf("1 + 0.2 = %v, %v\n", a, FormatFloat(a, 10))

	b := Subtract(0.3, 0.1)
	fmt.Printf("0.3 - 0.1 = %v, %v\n", b, FormatFloat(b, 10))

	c := Multiply(1.005, 100)
	fmt.Printf("1.005 * 100 = %v, %v\n", c, FormatFloat(c, 10))

	d, err := Divide(0.3, 0.1)
	if err == nil {
		fmt.Printf("0.3 / 0.1 = %v, %v\n", d, FormatFloat(d, 10))
	}

	e := Round(1.2345, 2)
	fmt.Printf("Round(1.2345, 2) = %v\n", e)

	_, err = Divide(1.0, 0.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
