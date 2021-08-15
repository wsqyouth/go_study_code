
package lru

import "container/list"

// Catche is a LRU catche. 非并发安全
type Cache struct {
    maxBytes int64 //最大字节数
    nbytes int64 //当前已使用
    ll *list.List
    cache map[string] *list.Element
    // 可选 当一个entry 被废弃时被回调
    OnEvicted func(key string, value Value)
}

type entry struct{
    key string
    val Value
}

// Value 使用Len（）计算已使用多少字节
type Value interface{
    Len() int
}

// New is the contructor of Cache
func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
    return &Cache{
        maxBytes:maxBytes,
        ll:list.New(),
        cache:make(map[string]*list.Element),
        OnEvicted:onEvicted,
    }
}


//Get look up a key's value
func (c *Cache) Get(key string) (value Value, ok bool) {
    if ele, ok := c.cache[key]; ok {
        c.ll.MoveToFront(ele)
        kv := ele.Value.(*entry)
        return kv.value,true
    }
    return
}

// RemoveOldest remove the oldest item
func (c *Cache) RemoveOldest(){
    ele := c.ll.Back()
    if ele != nil {
        c.ll.Remove(ele)
        kv := ele.Value.(*entry)
        delete(c.cache,kv.key) //从字典中删除该节点映射关系
        c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len()) //更新所使用内存
        if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
    }
}

// Add add a value to the catche
func (c *Cache) Add(key string ,value Value){
    if ele,ok := c.cache[key]; ok {
        c.ll.MoveToFront[ele]
        kv := ele.Value.(*entry)
        c.nbytes += int64(value.Len()) - int64(kv.value.Len())
        kv.value = value
    }else{
        ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
    }

    for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

// Len the number of cache entries
func (c *Cache) Len() int {
	return c.ll.Len()
}
