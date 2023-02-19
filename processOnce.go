package main

import (
	"context"
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx := initOnceObj(context.Background())
	obj := ProcessObj{
		id:   1,
		name: "paopao",
	}
	g, ctx := errgroup.WithContext(ctx)
	for i := 0; i < 10; i++ {
		g.Go(func() error {
			if err := processObj(ctx, obj); err != nil {
				return err
			}
			return nil
		})
	}
	err := g.Wait()
	if err != nil {
		panic("panic")
	}
}

// processObj 处理对象
func processObj(ctx context.Context, obj ProcessObj) error {
	onceObj, ok := getOnceObj(ctx)
	if ok && !onceObj.SetNxObj(ctx, obj) {
		fmt.Printf("ignore concurrent pocessObj: %v\n", obj)
		return nil
	}
	fmt.Printf("process obj: id:%v,name: %v\n", obj.id, obj.name)
	return nil
}

type OnceObjType int

const (
	OnceObjTypeKey OnceObjType = 1
)

// initOnceObj 初始化ctx key
func initOnceObj(ctx context.Context) context.Context {
	return context.WithValue(ctx, OnceObjTypeKey, new(OnceObj))
}

// getOnceObj 从ctx获取key
func getOnceObj(ctx context.Context) (*OnceObj, bool) {
	obj, ok := ctx.Value(OnceObjTypeKey).(*OnceObj)
	return obj, ok
}

type ProcessObj struct {
	name string
	id   int64
}

type OnceObj struct {
	mutex         sync.Mutex
	processObjMap map[int64]ProcessObj
}

func genProcessObjKey(obj ProcessObj) int64 {
	return obj.id
}

// SetNxObj 仅处理一次obj,如已被处理过则不再处理
func (o *OnceObj) SetNxObj(ctx context.Context, processObj ProcessObj) bool {
	o.mutex.Lock()
	defer o.mutex.Unlock()

	if o.processObjMap == nil {
		o.processObjMap = make(map[int64]ProcessObj)
	}

	key := genProcessObjKey(processObj)
	_, exist := o.processObjMap[key]
	if !exist {
		o.processObjMap[key] = processObj
	}
	// 未被处理过
	return !exist
}

//学习心得:
/*
在没有分布式锁时，如何自己实现一个实现单机请求只处理一次
这里借助map和ctx的特性,利用并发库进行实战测试，思想值得学习
*/
