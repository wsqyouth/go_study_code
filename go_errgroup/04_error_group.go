package main

import (
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"golang.org/x/sync/errgroup"
)

//errgroup包在sync.WaitGroup功能的基础上，增加了错误传递，以及在发生不可恢复的错误时取消整个goroutine集合，或者等待超时
func main() {

	var g errgroup.Group

	g.Go(func() error {
		fmt.Println("Goroutine 1")
		return nil //返回nil
	})
	g.Go(func() error {
		fmt.Println("Goroutine 2")
		return errors.New("Goroutine error") //返回自定义错误
	})
	if err := g.Wait(); err != nil {
		fmt.Println("Get errors: ", err)
	} else {
		fmt.Println("Success")
	}

	var urls = []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.qq.cm",
	}
	err := batchGetResult(context.Background(), urls)
	if err != nil {
		fmt.Printf("batchGetURLResult err. urls: %v\n", urls)
	}
}

//

func batchGetResult(ctx context.Context, urls []string) (err error) {
	g, ctx := errgroup.WithContext(ctx)
	for _, url := range urls {

		url := url
		g.Go(func() error {

			//获取url body
			//获取http请求并打印返回码
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println(resp.Status)
				resp.Body.Close()
			}
			return err
		})
	}
	//等待所有请求结束
	if err := g.Wait(); err == nil {
		fmt.Println("Succ fetch all urls")
	}
	return nil
}
