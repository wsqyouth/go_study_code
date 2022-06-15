package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// 动机：使用go routine和chan实现多对1的pipeline,即fan in/out
// demo: 通过并发的方式对一个很长的数组中的质数进行求和：先把数组分段,然后集中求和[类似map reduce]
func main() {
	nums := makeRange(1, 10)
	in := echo(nums)

	const nProcess = 5
	var chans [nProcess]<-chan int
	for i := range chans {
		chans[i] = sum(prime(in))
	}
	for n := range sum(merge(chans[:])) {
		fmt.Println(n)
	}
	fmt.Println("nvim-go")
	time.Sleep(time.Second)
}

// makeRange 生成1-10000的数组
func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

// echo 把数组全部echo到一个channel里[疑问:这种写法不会内存泄漏? 如果go routine没执行完主线程就return了呢]
func echo(nums []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func is_prime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

//prime 找出chan里的所有质数
func prime(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			if is_prime(n) {
				out <- n
			}
		}
		close(out)
	}()
	return out
}

// merge 把所有子数组的结果放到一个数组里
func merge(cs []<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	wg.Add(len(cs))
	for _, c := range cs {
		go func(c <-chan int) {
			for n := range c {
				out <- n
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// sum 求和拼接起来,得到最终的结果
func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum = 0
		for n := range in {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}

//-----
//总结：coolshell把go的特性用的出神入化，值得学习。本篇文章也是为pipleline的后续学习打下基础
//ref: https://coolshell.cn/articles/21228.html
