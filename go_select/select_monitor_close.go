package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker struct {
	id string
}

func (w Worker) Close(ch chan struct{}) error {
	fmt.Println("close: ", w.id)
	ch <- struct{}{}
	return nil
}

type Monitor struct {
	workers map[string]Worker // 监管每一个工人
}

// Close 通知各个woker执行关闭,最长等待10s
func (m *Monitor) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	var wg sync.WaitGroup
	for _, worker := range m.workers {
		wg.Add(1)
		go func(w Worker) {
			defer wg.Done()
			c := make(chan struct{}, 1)
			go w.Close(c)

			select {
			case <-c:
			case <-ctx.Done():
			}
		}(worker)
	}
	fmt.Println("monitor continuing...")
	wg.Wait()
	//继续观察看是否go routine是否仍在执行,如果没有监控输出，就表示停止了
	fmt.Println("all finish...")
	return nil
}
func main() {
	workers := make(map[string]Worker)
	workers["zhangsan"] = Worker{id: "111"}
	workers["lisi"] = Worker{id: "222"}
	var m Monitor
	m.workers = workers
	m.Close()
}

//文档参考：
// https://github.com/iswbm/GolangCodingTime/blob/master/source/c04/c04_08.md
