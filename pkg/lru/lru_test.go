package lru

import (
	"fmt"
	"reflect"
	"testing"
)

type String string

func (d String) Len() int {
	return len(d)
}

// TestCache_Get：测试LRU缓存池的获取方法
func TestCache_Get(t *testing.T) {

	lru := New(int64(0), nil)
	lru.Add("key1", String("123123123"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "123123123" {
		t.Fatalf("cache hit key1 = 123123123 failed")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}

}

// TestCache_Add：测试LRU缓存池的添加方法
func TestCache_Add(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key", String("1"))
	lru.Add("key", String("111"))

	if lru.nowData != int64(len("key")+len("111")) {
		t.Fatal("expected nowData is 6 but got : ", lru.nowData)
	}
}

// TestCache_RemoveOldest：测试LRU缓存池的移除最少访问节点
func TestCache_RemoveOldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"

	cap := len(k1 + k2 + v1 + v2) // 设置容量
	lru := New(int64(cap), nil)   // 设置一个最大容量为cap的Cache
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3)) // 此时会将k1挤出去

	// 测试k1能否获取，如果能获取，证明移除最少访问的节点失败
	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("RemoveOldest key1 failed")
	}
}

// TestOnEvicted：测试LRU缓存池的回调函数
// 某条记录被移除时会调用回调函数，返回被删除的key 和 value
func TestOnEvicted(t *testing.T) {

	keys := make([]string, 0)
	// 进行回调方法的定义，这里不断追加被删除节点的key到keys中
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}

	lru := New(int64(10), callback) // New Cache的时候传入callback回调函数

	lru.Add("key1", String("123456"))
	fmt.Printf("the cap of key:key1 + string:123456 is %d \n", int64(len("key1"))+int64(len("123456")))
	lru.Add("k2", String("k2")) // 当前内存超过最大限定，此时会调用回调方法，返回key1及其value
	lru.Add("k3", String("k3"))
	fmt.Printf("the cap of key:k3 + string:k3 is %d \n", int64(len("k3"))+int64(len("k3")))
	lru.Add("k4", String("k4")) // 当前内存超过最大限定，此时会再次调用回调方法，返回key2及其value

	// 打印当前keys
	for _, value := range keys {
		fmt.Printf(" %s\n", value)
	}

	expect := []string{"key1", "k2"} // 期望获得的数据,看是否匹配

	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys equals to %s", expect)
	}
}
