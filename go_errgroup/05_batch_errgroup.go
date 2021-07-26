package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/context"

	"golang.org/x/sync/errgroup"
)

const defultMaxGoRoutine = 2 //最大可设置并发数量

//errgroup包在sync.WaitGroup功能的基础上，增加了错误传递，以及在发生不可恢复的错误时取消整个goroutine集合，或者等待超时
func main() {
	var urls = []string{
		"http://www.baidu.com",
		"http://www.google.com",
		"http://www.qq.cm",
	}
	err := batchGetResultNew(context.Background(), urls)
	if err != nil {
		fmt.Printf("batchGetURLResult err. urls: %v\n", urls)
	}
}

func batchGetResultNew(ctx context.Context, aidList []string) (err error) {

	nLen := len(aidList)
	maxGoroutine := nLen
	if nLen > defultMaxGoRoutine {
		maxGoroutine = defultMaxGoRoutine
	}
	ch := make(chan string)
	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i < maxGoroutine; i++ {
		g.Go(func() (e error) {
			for each := range ch {
				//获取url body
				//获取http请求并打印返回码
				resp, err := http.Get(each)
				if err == nil {
					fmt.Println(resp.Status)
					resp.Body.Close()
				}
				return err
			}

			return nil
		})
	}

	for _, each := range aidList {
		ch <- each
	}

	// close 后前面的for循环才能结束
	close(ch)
	//等待所有请求结束
	if err := g.Wait(); err != nil {
		return err
	}
	fmt.Println("Succ fetch all urls")
	return nil
}
