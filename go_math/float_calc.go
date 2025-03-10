package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// Add 基础加法运算（常规浮点实现）
func Add(a, b float64) float64 {
	return a + b
}

// Subtract 基础减法运算
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply 基础乘法运算
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide 基础除法运算
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// Round 四舍五入实现（处理浮点精度问题）
func Round(num float64, places int) float64 {
	if places < 0 {
		places = 0
	}

	pow := math.Pow10(places)
	value := num * pow
	_, frac := math.Modf(value + math.Copysign(0.5, value))

	return (value - frac) / pow
}

// FormatFloat 保持原样（与之前实现相同）
func FormatFloat(num float64, places int) string {
	if places < 0 {
		places = 0
	}
	return strconv.FormatFloat(num, 'f', places, 64)
}

func main() {
	// 测试用例保持不变
	a := Add(0.1, 0.2)
	fmt.Printf("0.1 + 0.2 = %v, %v\n", a, FormatFloat(a, 10))

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
