package main

import (
	"fmt"
	"time"
)

//trace 编译
func main() {

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello gmp")
	}
}

// 测试:通过go tool trace 打开trace文件进行分析
// go build 01_trace.go
// GODEBUG=schedtrace=1000 ./01_trac
// SCHED 0ms: gomaxprocs=2[当前多少个p] idleprocs=0[空闲] threads=4[多少个m] spinningthreads=1 idlethreads=1[空闲] runqueue=0[全局队列中G数量] [2 0] // 2个p对应本地G数量的个数
// linux: top 1, 可以看到当前机器2
