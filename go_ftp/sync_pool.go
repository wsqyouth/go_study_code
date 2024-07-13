package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var p sync.Pool // 创建一个对象池
	p.New = func() interface{} {
		return &http.Client{
			Timeout: 5 * time.Second,
		}
	}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			client := p.Get().(*http.Client)
			defer p.Put(client)
			//获取http请求并打印返回码
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
