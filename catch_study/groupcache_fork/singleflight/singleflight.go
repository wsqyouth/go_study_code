/*
Copyright 2012 Google Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

//这主要是进行相同访问的一个合并操作。也就是说，如果对于某个key的请求已经存在并且正在进行，则对该key的新的请求会堵塞在这里，
//等原来的请求结束后，将请求得到的结果同时返回给堵塞中的请求

// Package singleflight provides a duplicate function call suppression
// mechanism.
package singleflight

import "sync"

// call is an in-flight or completed Do call  实际请求的封装
type call struct {
	wg  sync.WaitGroup
	val interface{} //实际的请求函数
	err error
}

// 主要是用来组织已经存在的对某key的请求和对应的实际请求函数映射
// Group represents a class of work and forms a namespace in which
// units of work can be executed with duplicate suppression.
type Group struct {
	mu sync.Mutex       // protects m
	m  map[string]*call // lazily initialized
}

// Do executes and returns the results of the given function, making
// sure that only one execution is in-flight for a given key at a
// time. If a duplicate comes in, the duplicate caller waits for the
// original to complete and receives the same results.
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m == nil {
		g.m = make(map[string]*call)
	}
	//如果m中存在对该key的请求，则该线程不会直接再次访问key，所以释放锁
	//然后堵塞等待已经存在的请求得到的结果
	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait() //牛逼
		return c.val, c.err
	}
	//如果不存在对该key的请求，则本线程要进行实际的请求，保持m的锁定状态
	c := new(call)  //创建一个实际请求结构体
	c.wg.Add(1) //为了保证其他的相同请求的堵塞
	g.m[key] = c
	g.mu.Unlock()

	//执行真正的请求函数，得到对该key请求的结果
	c.val, c.err = fn()
	c.wg.Done() //得到结果后取消其他请求的堵塞

	g.mu.Lock()
	delete(g.m, key)//该次请求完成后，要从已存在请求map中删掉
	g.mu.Unlock()

	return c.val, c.err //返回请求结果
}
