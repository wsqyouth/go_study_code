package main

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

func main() {
	// 创建一个新的缓存
	cache := gcache.New(20).
		LRU().
		Expiration(time.Minute).
		LoaderFunc(func(key interface{}) (interface{}, error) {
			// 加载函数，当缓存中不存在某个键的值时，会调用该函数加载该键的值
			fmt.Println("loading key:", key)

			// 执行相关操作，加载该键对应的值，这里使用一个简单的示例
			value := fmt.Sprintf("value of key: %v", key)

			// 返回加载得到的值，并且返回error
			return value, nil
		}).
		Build()

	// 获取键"foo"的值
	value, err := cache.Get("foo")
	if err != nil {
		fmt.Println("get value error:", err)
	} else {
		fmt.Println("get value:", value)
	}

	// 再次获取键"foo"的值，此时值已经被缓存起来了，不会再次调用LoaderFunc函数
	value, err = cache.Get("foo")
	if err != nil {
		fmt.Println("get value error:", err)
	} else {
		fmt.Println("get value:", value)
	}
}

/*
可以看到，第一次获取键"foo"的值时，调用了加载函数，并成功地加载了值。
而第二次获取键"foo"的值时，直接从缓存中获取到了之前加载的值，没有再调用加载函数。
*/
