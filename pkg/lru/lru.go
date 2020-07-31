package lru

import (
	"container/list"
)

// 创建结构体 Cache，方便实现后续的增改删查工作
type Cache struct {
	maxData   int64                         // 允许使用最大内存
	nowData   int64                         // 当前已使用内存
	list      *list.List                    // LRU底层数据结构：双向链表
	cache     map[string]*list.Element      // 键是字符串，值是双向链表中对应节点的指针
	OnEvicted func(key string, value Value) // 某条记录被移除时的回调函数
}

// entry：双向链表的节点的数据类型
// 所以我们这次实现的LRU底层结构是双向链表
type entry struct {
	// 在链表中仍保存每个值对应的 key 的好处在于：淘汰队首节点时，需要用 key 从字典中删除对应的映射
	key   string
	value Value
}

// Value：实现Value 接口的任意类型
type Value interface {
	// 接口只包含了一个方法 Len() int，用于返回值所占用的内存大小
	Len() int
}

// New：方便实例化Cache，实现New函数
// 传入参数有maxData，和一个onEvicted 方法
func New(maxData int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxData:   maxData,
		list:      list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

// Get：实现获取功能
// 根据key进行value的查找，并返回一个是否查找成功的标志
// 根据LRU，将查找之后的节点移动到维护的双向队列的队尾（最久没有访问的放到队头）
func (c *Cache) Get(key string) (value Value, ok bool) {
	// 1.第一步是从字典中找到对应的双向链表的节点
	if ele, ok := c.cache[key]; ok {
		// 2.将链表中的节点 ele 移动到队尾，这里约定front作为队尾
		c.list.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest ：实现删除功能
// 根据LRU，移除最近最少访问节点即可，即移除队首元素即可
func (c *Cache) RemoveOldest() {
	ele := c.list.Back() // 取队首节点
	if ele != nil {
		c.list.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)                                 // 从map中删除该节点的映射关系
		c.nowData -= int64(len(kv.key)) + int64(kv.value.Len()) // 更新内存
		if c.OnEvicted != nil {                                 // 若回调函数不为nil，则调用回调函数
			c.OnEvicted(kv.key, kv.value)
		}
	}
}

// 实现新增/修改功能
func (c *Cache) Add(key string, value Value) {
	// 1.1 如果key在Map中存在，则更新对应节点的值，并将该节点移到队尾。
	if ele, ok := c.cache[key]; ok {
		c.list.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nowData += int64(value.Len()) - int64(kv.value.Len()) // 更新内存
		kv.value = value
	} else {
		// 1.2 如果key在Map不存在，则向队尾进行添加新节点，并在Map中添加映射关系
		ele := c.list.PushFront(&entry{key, value})       // 添加新节点
		c.cache[key] = ele                                // 添加Map映射关系
		c.nowData += int64(len(key)) + int64(value.Len()) // 更新内存
	}

	// 2.更新 c.nbytes，如果超过了设定的最大值 c.maxBytes，则移除最少访问的节点。
	for c.maxData != 0 && c.nowData > c.maxData {
		c.RemoveOldest()
	}
}

// 获取Cache添加了多少条数据
func (c *Cache) Len() int {
	return c.list.Len()
}
