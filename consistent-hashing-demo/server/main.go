package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type CachedMap struct {
	KvMap sync.Map
	Lock  sync.RWMutex
}

var (
	cache = CachedMap{KvMap: sync.Map{}}

	port = flag.String("p", "8080", "port")

	regHost = "http://localhost:18888"

	expireTime = 10
)

/*
go run server/main.go -p 8081
*/
func main() {
	flag.Parse()

	stopChan := make(chan interface{})
	startServer(*port)
	<-stopChan
}

/*
缓存服务器启动
开启几台缓存服务器,通过registerHost进行注册
使用sync.Map模拟缓存,key存在则返回,不存在则写入.
通过time.AfterFunc设置缓存时间
*/
func startServer(port string) {
	hostName := fmt.Sprintf("localhost:%s", port)

	fmt.Printf("start server: %s\n", port)

	err := registerHost(hostName)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", kvHandle)
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		err = unregisterHost(hostName)
		if err != nil {
			panic(err)
		}
		panic(err)
	}
}

/*
核心在于每个缓存服务器有自己的map进行存取
*/
func kvHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	if _, ok := cache.KvMap.Load(r.Form["key"][0]); !ok {
		val := fmt.Sprintf("hello: %s", r.Form["key"][0])
		cache.KvMap.Store(r.Form["key"][0], val)
		fmt.Printf("cached key: {%s: %s}\n", r.Form["key"][0], val)

		time.AfterFunc(time.Duration(expireTime)*time.Second, func() {
			cache.KvMap.Delete(r.Form["key"][0])
			fmt.Printf("removed cached key after 10s: {%s: %s}\n", r.Form["key"][0], val)
		})
	}

	val, _ := cache.KvMap.Load(r.Form["key"][0])

	_, err := fmt.Fprintf(w, val.(string))
	if err != nil {
		panic(err)
	}
}

/*
类似服务发现,向代理服务器注册自身
*/
func registerHost(host string) error {
	resp, err := http.Get(fmt.Sprintf("%s/register?host=%s", regHost, host))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

/*
类似服务发现,向代理服务器注销自身
*/
func unregisterHost(host string) error {
	resp, err := http.Get(fmt.Sprintf("%s/unregister?host=%s", regHost, host))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
