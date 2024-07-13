package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var p sync.Pool // 创建一个对象池
	for i := 0; i < 5; i++ {
		p.Put(&http.Client{Timeout: 5 * time.Second}) // 不设置 New 字段,初始化时就放入5个可重用对象
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			client, ok := p.Get().(*http.Client)
			if !ok {
				fmt.Println("get client is nil")
				return
			}
			defer p.Put(client)
			resp, err := client.Get("https://www.baidu.com")
			if err != nil {
				fmt.Println("http get error", err)
				return
			}
			resp.Body.Close()
			fmt.Println("success", resp.Status)
		}()
	}
	//等待所有请求结束
	wg.Wait()
}
