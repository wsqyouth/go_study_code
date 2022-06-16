package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// 装饰器模式
func timedSumFunc(f SumFunc) SumFunc {
	return func(start, end int64) int64 {
		defer func(t time.Time) {
			fmt.Printf("--- Time Elapsed(%s): %v ---\n", getFunctionName(f), time.Since(t))
		}(time.Now())
		return f(start, end)
	}
}

func SumOld(start, end int64) int64 {
	var sum int64
	sum = 0
	if start > end {
		start, end = end, start
	}
	for i := start; i <= end; i++ {
		sum += i
	}
	return sum
}
func SumNew(start, end int64) int64 {
	if start > end {
		start, end = end, start
	}
	return (end - start + 1) * (end + start) / 2
}

// 统计两个求和函数的性能优劣
func main() {
	sumDecoraOld := timedSumFunc(SumOld)
	sumDecoraNew := timedSumFunc(SumNew)
	fmt.Printf("%d, %d\n", sumDecoraOld(-10000, 1000000), sumDecoraNew(-10000, 1000000))
}

//ref: 后续扩展：1）泛型支持通用场景  2) 多个装饰圈可以使用pipeline进行包装 decoratePipeLine(targetFunc, a, b, c)
// https://coolshell.cn/articles/17929.html
