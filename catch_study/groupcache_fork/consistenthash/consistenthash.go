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

// Package consistenthash provides an implementation of a ring hash.
package consistenthash

import (
	"hash/crc32"
	"sort"
	"strconv"
)

//首先定义了一个函数Hash，用于将key值 Hash成32位整数，
//然后定义了Map结构体用于存放Hash后的结果，其中
//hash是上面的hash函数，Map结构中replicas的含义是增加虚拟桶，使数据分布更加均匀
//keys存放hash后的结果，并且经过了排序，其实就是一致性hash圆环，
//hashMap就是存放具体的对应，将key对应上hash后的32位整数。

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash
	replicas int
	keys     []int // Sorted
	hashMap  map[int]string
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE //注意如果hash函数为nil，则默认Hash函数为crc32库的ChecksumIEEE函数
	}
	return m
}

// IsEmpty returns true if there are no items available.
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		//将i与key拼接之后再进行Hash
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			//keys数组增加key
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	//先找到key对应的Hash值
	hash := int(m.hash([]byte(key)))

	// Binary search for appropriate replica. 二分查找key对应hash值之后的最近的一个节点
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	// Means we have cycled back to the first replica.
	//如果idx是keys的最后一个元素，则定向到第一个，因为模拟的是环
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}


//一致性算法讲解：数据倾斜使用虚拟节点解决
https://geektutu.com/post/geecache-day4.html
https://www.cyhone.com/articles/consistent-hash-of-groupcache/ [好文]
//实例算法讲解：https://blog.csdn.net/mrbuffoon/article/details/83584052