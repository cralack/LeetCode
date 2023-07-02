package lrucache

import (
	"testing"
)

type LRUCache struct {
	Size, Cap  int
	Head, Tail *DLNode
	Cache      map[int]*DLNode
}
type DLNode struct {
	Key, Val   int
	Next, Prev *DLNode
}

func Constructor(capacity int) LRUCache {
	// 尾部的数据是最近使用的，头部的数据是最久为使用的
	head := &DLNode{Key: -1}
	tail := &DLNode{Key: -2}
	head.Next = tail
	tail.Prev = head
	cache := make(map[int]*DLNode)
	lru := &LRUCache{
		Size:  0,
		Cap:   capacity,
		Head:  head,
		Tail:  tail,
		Cache: cache,
	}

	return *lru
}
func (lr *LRUCache) Remove(x *DLNode) int {
	key := x.Key
	x.Next.Prev = x.Prev
	x.Prev.Next = x.Next
	lr.Size--
	return key
}
func (lr *LRUCache) RemoveFirst() int {
	if lr.Head.Next == lr.Tail {
		return -1
	}
	first := lr.Head.Next
	return lr.Remove(first)
}
func (lr *LRUCache) AddLast(x *DLNode) {
	x.Prev = lr.Tail.Prev
	x.Next = lr.Tail
	lr.Tail.Prev.Next = x
	lr.Tail.Prev = x
	lr.Size++
}
func (lr *LRUCache) MakeRecently(x *DLNode) {
	x.Next.Prev = x.Prev
	x.Prev.Next = x.Next
	x.Prev = lr.Tail.Prev
	x.Next = lr.Tail
	lr.Tail.Prev.Next = x
	lr.Tail.Prev = x
}

func (lr *LRUCache) Get(key int) int {
	if value, ok := lr.Cache[key]; ok {
		// 提升优先级
		lr.MakeRecently(value)
		return value.Val
	}
	return -1
}

func (lr *LRUCache) Put(key int, value int) {
	if v, ok := lr.Cache[key]; ok {
		v.Val = value
		// 提升优先级
		lr.MakeRecently(v)

	} else {
		// cache满了
		tmp := &DLNode{Key: key, Val: value}
		if lr.Size == lr.Cap {
			rKey := lr.RemoveFirst() // map need --
			delete(lr.Cache, rKey)
		} // else { //cache 没满
		// }
		lr.AddLast(tmp)
		lr.Cache[key] = tmp
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func Test_LRU_cache(t *testing.T) {
	command1 := []string{
		"LRUCache", "put", "put",
		"get", "put", "get", "put",
		"get", "get", "get"}
	conmand2 := [][]int{
		{2}, {1, 1}, {2, 2},
		{1}, {3, 3}, {2}, {4, 4},
		{1}, {3}, {4}}
	var lru LRUCache
	for i, v := range command1 {
		switch v {
		case "LRUCache":
			lru = Constructor(conmand2[i][0])
		case "put":
			lru.Put(conmand2[i][0], conmand2[i][1])
		case "get":
			t.Log(lru.Get(conmand2[i][0]))
		}
	}
}
