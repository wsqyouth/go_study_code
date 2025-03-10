package main

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
)

// 设置精度级别
const precision = 64

// Add 精确加法运算
func Add(a, b float64) float64 {
	bigA := new(big.Float).SetPrec(precision).SetFloat64(a)
	bigB := new(big.Float).SetPrec(precision).SetFloat64(b)

	result := new(big.Float).Add(bigA, bigB)
	val, _ := result.Float64()
	return val
}

// Subtract 精确减法运算
func Subtract(a, b float64) float64 {
	bigA := new(big.Float).SetPrec(precision).SetFloat64(a)
	bigB := new(big.Float).SetPrec(precision).SetFloat64(b)

	result := new(big.Float).Sub(bigA, bigB)
	val, _ := result.Float64()
	return val
}

// Multiply 精确乘法运算
func Multiply(a, b float64) float64 {
	bigA := new(big.Float).SetPrec(precision).SetFloat64(a)
	bigB := new(big.Float).SetPrec(precision).SetFloat64(b)

	result := new(big.Float).Mul(bigA, bigB)
	val, _ := result.Float64()
	return val
}

// Divide 精确除法运算
func Divide(a, b float64) (float64, error) {
	// 检查除数是否为零
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	bigA := new(big.Float).SetPrec(precision).SetFloat64(a)
	bigB := new(big.Float).SetPrec(precision).SetFloat64(b)

	result := new(big.Float).Quo(bigA, bigB)
	val, _ := result.Float64()
	return val, nil
}

// Round 四舍五入到指定小数位
func Round(num float64, places int) float64 {
	// 确保 places 为非负数
	if places < 0 {
		places = 0
	}

	bigNum := new(big.Float).SetPrec(precision).SetFloat64(num)

	// 创建表示小数位的缩放因子
	scale := new(big.Float).SetInt(new(big.Int).Exp(
		big.NewInt(10),
		big.NewInt(int64(places)),
		nil,
	))

	// 乘以缩放因子
	bigNum.Mul(bigNum, scale)

	// 四舍五入到整数
	rounded := new(big.Float)
	rounded.Add(bigNum, big.NewFloat(0.5))

	// 截断小数部分
	intPart := new(big.Int)
	rounded.Int(intPart)

	// 将结果除以缩放因子
	result := new(big.Float).SetInt(intPart)
	result.Quo(result, scale)

	val, _ := result.Float64()
	return val
}

// FormatFloat 将浮点数格式化为指定小数位的字符串，避免科学计数法
func FormatFloat(num float64, places int) string {
	if places < 0 {
		places = 0
	}
	return strconv.FormatFloat(num, 'f', places, 64)
}

func main() {
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

	// 处理四舍五入
	e := Round(1.2345, 2)
	fmt.Printf("Round(1.2345, 2) = %v\n", e, FormatFloat(e, 2))

	// 测试错误处理
	_, err = Divide(1.0, 0.0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
