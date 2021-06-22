package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"log"
)

var (
	rdb *redis.Client
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// 使用Ping方法测试连接是否通畅
	pong, err := client.Ping().Result()
	if err != nil {
		log.Printf("redis连接失败,错误信息:%v\n", err)
		return
	}
	log.Printf("redis连接成功. %v", pong)

	// Set
	err = client.Set("name", "coopers", 0).Err()
	if err != nil {
		panic(err)
	}

	// Get
	str := client.Get("name")
	log.Println(str)

	//添加键值对
	err = client.Set("golang", "yes", 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("键golang设置成功")

	// 判断 key 是否存在
	val, err := client.Get("name").Result()
	if err == redis.Nil {
		fmt.Println("name  does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("name: ", val)
	}

	V7Example()
}

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",  // no password set
		DB:       0,   // use default DB
		PoolSize: 100, // 连接池大小
	})

	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()

	_, err = rdb.Ping().Result()
	return err
}

func V7Example() {
	if err := initClient(); err != nil {
		return
	}

	err := rdb.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get("key").Result()
	if err != nil {
		panic(err)
	}
    fmt.Println("key:", val)

	val2, err := rdb.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}
