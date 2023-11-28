package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func generateNumbers() []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, 10)
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Intn(100) + 1
	}
	return numbers
}

func processNumber(ctx context.Context, num int) (int, error) {
	// 这里模拟处理流水的过程，我们简单地返回输入的平方
	return num * num, nil
}

// 添加一个函数preprocessNumbers，这个函数会检查输入的数字是否都是正数。如果有负数，函数就返回一个错误。
func preprocessNumbers(numbers []int) error {
	for _, num := range numbers {
		if num < 0 {
			return fmt.Errorf("negative number found: %d", num)
		}
	}
	return nil
}

func main() {

	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()
	for {
		numbers := generateNumbers()      // 这是一个假设的函数，用于生成新的数字
		numbers = append(numbers, -1)     // 这是我们的流水数据,error,此时可能会有泄漏
		doneNumbers := make(map[int]bool) // 用于记录已经处理过的数据

		ch := make(chan int)
		maxGoroutine := len(numbers)
		if maxGoroutine > 128 {
			maxGoroutine = 128
		}

		// 在channel上等待的goroutine可能会永远等待，这就是所谓的goroutine泄漏。
		var g sync.WaitGroup
		for i := 0; i < maxGoroutine; i++ {
			g.Add(1)
			go func() {
				defer g.Done()
				for num := range ch {
					res, err := processNumber(context.Background(), num)
					if err != nil {
						fmt.Println("Error processing number:", err)
						return
					}
					fmt.Printf("processNumber num:%v, res:%v\n", num, res)
				}
			}()
		}

		err := preprocessNumbers(numbers)
		if err != nil {
			close(ch) // 这里出错未关闭channel直接返回时,可能会有泄漏
			fmt.Println("Error in preprocessing:", err)
			continue
		}

		// 将未处理的数据发送到channel
		for _, num := range numbers {
			if _, ok := doneNumbers[num]; ok {
				continue
			}
			ch <- num
		}

		close(ch) // 关闭channel，使得goroutine能够结束
		g.Wait()  // 等待所有goroutine结束

		fmt.Println("All numbers processed.")
	}
}

/*
在项目中preprocessNumbers未考虑close channel，导致内存泄漏
go tool pprof http://localhost:6060/debug/pprof/goroutine?debug=2
排查：检查goroutin为13w, 并且很多go routine在channel等待recv
*/
