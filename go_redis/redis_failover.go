package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func main() {
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		SentinelAddrs: []string{"127.0.0.1:26377"}, // 哨兵节点地址
		MasterName:    "master-redis",              // 主节点名称
		Password:      "",
		DB:            0,
	})

	ctx := context.Background()
	result := client.Ping(ctx)
	fmt.Println(result.Val())
}
