package safemap

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Map结构体
type Map struct {
	sm    sync.Map
	count int64
}

// New ： 构建一个新的map
func New() *Map {
	return &Map{
		count: 0,
	}
}

// Load ： 根据key去map中进行查询
func (m *Map) Load(key string) (value interface{}, ok bool) {
	return m.sm.Load(key)
}

// MustLoad ：根据key去map中进行硬查询
func (m *Map) MustLoad(key string) interface{} {
	v, ok := m.Load(key)
	if !ok {
		panic(fmt.Errorf("key %s not exist", key))
	}
	return v
}

// Range ：范围查询
func (m *Map) Range(f func(key string, value interface{}) bool) {
	m.sm.Range(func(key, value interface{}) bool {
		return f(key.(string), value)
	})
}

// Store ：将键值对存入map，并对计数count+1
func (m *Map) Store(key string, value interface{}) {
	atomic.AddInt64(&m.count, 1)
	m.sm.Store(key, value)
}

// Delete : 根据key进行删除，并对计数count-1
func (m *Map) Delete(key string) {
	if m.count > 0 {
		atomic.AddInt64(&m.count, -1)
	}
	m.sm.Delete(key)
}

// Length ：返回map的长度
func (m *Map) Length() int64 {
	return m.count
}

// Empty ：判断是否为空
func (m *Map) Empty() bool {
	return m.count == 0
}
