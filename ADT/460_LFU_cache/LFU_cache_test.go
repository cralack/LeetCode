package lfucache

import "testing"

type Node struct {
	Key, Val   int
	Freq       int
	Next, Prev *Node
}
type LFUCache struct {
	Size, Cap int
	MinFreq   int
	KeyToNode map[int]*Node
	Cache     map[int]*DoubleList
	Cnt       map[int]int
}
type DoubleList struct {
	Head, Tail *Node
}

func Constructor(capacity int) LFUCache {
	lfu := &LFUCache{
		Size:      0,
		Cap:       capacity,
		MinFreq:   999,
		KeyToNode: make(map[int]*Node),
		Cache:     make(map[int]*DoubleList),
		Cnt:       make(map[int]int),
	}
	return *lfu
}
func (lfu *LFUCache) Add(x *Node, freq int) {
	if _, ok := lfu.Cache[freq]; !ok {
		head := &Node{Key: -1}
		tail := &Node{Key: -2}
		head.Next = tail
		tail.Prev = head
		lfu.Cache[freq] = &DoubleList{Head: head, Tail: tail}
	}
	x.Prev = lfu.Cache[freq].Head
	x.Next = lfu.Cache[freq].Head.Next
	lfu.Cache[freq].Head.Next.Prev = x
	lfu.Cache[freq].Head.Next = x
	lfu.Size++
	lfu.Cnt[freq]++
}
func (lfu *LFUCache) Remove(x *Node) {
	x.Next.Prev = x.Prev
	x.Prev.Next = x.Next
	x.Next = nil
	x.Prev = nil
	lfu.Size--
	lfu.Cnt[x.Freq]--
	if x.Freq == lfu.MinFreq && lfu.Cnt[lfu.MinFreq] == 0 {
		lfu.MinFreq++
	}
}

func (lfu *LFUCache) RemoveLast() *Node {
	x := lfu.Cache[lfu.MinFreq].Tail.Prev
	lfu.Remove(x)
	return x
}

func (lfu *LFUCache) IncreaseFreq(x *Node) {
	lfu.Remove(x)

	x.Freq++
	lfu.Add(x, x.Freq)
}

func (lfu *LFUCache) Get(key int) int {
	if lfu.Cap <= 0 {
		return -1
	}
	if v, ok := lfu.KeyToNode[key]; ok {
		lfu.IncreaseFreq(v)
		return v.Val
	}
	return -1
}

func (lfu *LFUCache) Put(key int, value int) {
	//already exist
	if v, ok := lfu.KeyToNode[key]; ok {
		v.Val = value
		lfu.IncreaseFreq(v)
		//del minfreq node
	} else if lfu.Cap == 0 {
		return
	} else {
		if lfu.Size >= lfu.Cap {
			min := lfu.RemoveLast()
			delete(lfu.KeyToNode, min.Key)
		}
		tmp := &Node{Key: key, Val: value, Freq: 1}
		lfu.Add(tmp, tmp.Freq)
		lfu.KeyToNode[key] = tmp
		lfu.MinFreq = 1
	}
}

/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func Test_LFU_cache(t *testing.T) {
	command1 := []string{
		"LFUCache", "put", "put",
		"get", "put", "get", "get",
		"put", "get", "get", "get"}
	command2 := [][]int{
		{2}, {1, 1}, {2, 2},
		{1}, {3, 3}, {2}, {3},
		{4, 4}, {1}, {3}, {4}}
	// command1 := []string{"LFUCache", "put", "get"}
	// command2 := [][]int{{0}, {0, 0}, {0}}
	var lfu LFUCache
	for i, v := range command1 {
		switch v {
		case "LFUCache":
			lfu = Constructor(command2[i][0])
		case "put":
			lfu.Put(command2[i][0], command2[i][1])
		case "get":
			t.Log(lfu.Get(command2[i][0]))
		}
	}
}
