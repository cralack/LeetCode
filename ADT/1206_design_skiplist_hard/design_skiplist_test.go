package designskiplisthard

import (
	"math/rand"
	"testing"
)

type SkipNode struct {
	Val  int
	Next []*SkipNode
}
type Skiplist struct {
	Level int
	Head  *SkipNode
}

func Constructor() Skiplist {
	return Skiplist{Level: 10,
		Head: &SkipNode{Val: -1, Next: make([]*SkipNode, 10)}}
}

func (this *Skiplist) Find(target int, Prevs []*SkipNode) {
	// 找出target在跳表中每一层的Prev节点并储存于Prevs数组中
	cur := this.Head
	for i := this.Level - 1; i >= 0; i-- { // 向下遍历
		for cur.Next[i] != nil && cur.Next[i].Val < target { // 严格<target
			cur = cur.Next[i]
		}
		Prevs[i] = cur
	} // 查找效率 O(logN)
}

func (this *Skiplist) Search(target int) bool {
	Prevs := make([]*SkipNode, this.Level)
	this.Find(target, Prevs)

	// 先判最底层 <target 的节点后继不为nil
	// 后判 val==target
	return Prevs[0].Next[0] != nil && Prevs[0].Next[0].Val == target
}

func (this *Skiplist) Add(target int) {
	Prevs := make([]*SkipNode, this.Level)
	this.Find(target, Prevs)

	node := &SkipNode{Val: target, Next: make([]*SkipNode, this.Level)}
	for i := 0; i < this.Level; i++ { // 由底向上
		node.Next[i] = Prevs[i].Next[i] // 连接target节点与跳表每层
		Prevs[i].Next[i] = node
		if rand.Intn(2) == 0 { // (1/2)^level概率被提取到上层
			break
		}
	}
}

func (this *Skiplist) Erase(target int) bool {
	Prevs := make([]*SkipNode, this.Level)
	this.Find(target, Prevs)

	node := Prevs[0].Next[0]               // 待删节点
	if node == nil || node.Val != target { // 判空
		return false
	}
	for i := 0; i < this.Level && Prevs[i].Next[i] == node; i++ {
		Prevs[i].Next[i] = node.Next[i] // 由下至上逐层断链
	}
	return true
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */
func Test_design_skiplist(t *testing.T) {
	c1 := []string{"Skiplist", "add", "add", "add", "search",
		"add", "search", "erase", "erase", "search"}
	c2 := []int{-1, 1, 2, 3, 0, 4, 1, 0, 1, 1}
	var sol Skiplist
	for i, c := range c1 {
		switch c {
		case "Skiplist":
			sol = Constructor()
		case "add":
			sol.Add(c2[i])
		case "search":
			t.Log(sol.Search(c2[i]))
		case "erase":
			t.Log(sol.Erase(c2[i]))
		}
	}
}
