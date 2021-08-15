/*
Copyright 2013 Google Inc.

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

//实现缓存的置换算法（最近最少使用）

// Package lru implements an LRU cache.
package lru

import "container/list"

// Cache is an LRU cache. It is not safe for concurrent access.
type Cache struct {
	// MaxEntries is the maximum number of cache entries before
	// an item is evicted. Zero means no limit.
	// MaxEntries是Cache中实体的最大数量，0表示没有限制。
	MaxEntries int

	// OnEvicted optionally specifies a callback function to be
	// executed when an entry is purged from the cache.
	//OnEvicted是一个回调函数，进行Cache操作达到一定条件可能需要回调做一些处理工作
	OnEvicted func(key Key, value interface{})

	ll    *list.List //ll是引用container/list包中的双向链表，是一个链表指针
	cache map[interface{}]*list.Element //cache是一个map，存放具体的k/v对，value是双向链表中的具体元素，也就是*Element
}

// A Key may be any value that is comparable. See http://golang.org/ref/spec#Comparison_operators
type Key interface{}

type entry struct {
	key   Key
	value interface{}
}

// New creates a new Cache.
// If maxEntries is zero, the cache has no limit and it's assumed
// that eviction is done by the caller.
func New(maxEntries int) *Cache {
	return &Cache{
		MaxEntries: maxEntries,
		ll:         list.New(),
		cache:      make(map[interface{}]*list.Element),
	}
}

// Add adds a value to the cache.[插入]
func (c *Cache) Add(key Key, value interface{}) {
	if c.cache == nil { //如果cache为空，重新初始化
		c.cache = make(map[interface{}]*list.Element)
		c.ll = list.New()
	}
	//如果cache中有该key，则该实体移动到头部（保证最近访问在最前）
	//然后将kv对实体中的value值替换
	if ee, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ee)
		ee.Value.(*entry).value = value
		return
	}
	//如果cache中没有该key，则将kv对封装成entry实体加入到双向链表头上
	ele := c.ll.PushFront(&entry{key, value})
	//让cache这个map的该key指向该kv对实体
	c.cache[key] = ele
	//检查Cache是否满，如果满，就移除最久未被访问的实体
	if c.MaxEntries != 0 && c.ll.Len() > c.MaxEntries {
		c.RemoveOldest()
	}
}

// Get looks up a key's value from the cache.[获取]
func (c *Cache) Get(key Key) (value interface{}, ok bool) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		// 如果该key存在则将其移到链表头部，因为最近被访问了
		c.ll.MoveToFront(ele)
		return ele.Value.(*entry).value, true
	}
	return
}

// Remove removes the provided key from the cache.[删除特定key]
func (c *Cache) Remove(key Key) {
	if c.cache == nil {
		return
	}
	if ele, hit := c.cache[key]; hit {
		c.removeElement(ele)
	}
}

// RemoveOldest removes the oldest item from the cache.[删除 oldest key]
func (c *Cache) RemoveOldest() {
	if c.cache == nil {
		return
	}
	ele := c.ll.Back()
	if ele != nil {
		c.removeElement(ele)
	}
}

func (c *Cache) removeElement(e *list.Element) {
	//双向链表删除该实体
	c.ll.Remove(e)
	kv := e.Value.(*entry)
	//同时清空相应map中的key
	delete(c.cache, kv.key)
	if c.OnEvicted != nil {
		c.OnEvicted(kv.key, kv.value)
	}
}

// Len returns the number of items in the cache. [获取当前Cache中的元素个数]
func (c *Cache) Len() int {
	if c.cache == nil {
		return 0
	}
	return c.ll.Len()
}

// Clear purges all stored items from the cache. [清空当前Cache]
func (c *Cache) Clear() {
	if c.OnEvicted != nil {
		for _, e := range c.cache {
			kv := e.Value.(*entry)
			c.OnEvicted(kv.key, kv.value)
		}
	}
	c.ll = nil
	c.cache = nil
}


//文章参考：
1. https://geektutu.com/post/geecache-day1.html   //参考图即可
2. https://blog.csdn.net/mrbuffoon/article/details/83547017 //算法讲解