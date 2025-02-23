package designskiplisthard

import (
	"math/rand"
	"testing"
	"time"
)

var rng = rand.New(rand.NewSource(time.Now().UnixNano()))

const (
	maxLevel = 16
	p        = 1 / 2.718281828459045
)

type node struct {
	val  int
	next []*node
}

func newNode(val, level int) *node {
	return &node{val: val, next: make([]*node, level)}
}

type Skiplist struct {
	head  *node
	level int
}

func Constructor() Skiplist {
	return Skiplist{head: newNode(-1, maxLevel), level: 1}
}

func (p *Skiplist) Search(target int) bool {
	cur := p.head
	for i := p.level - 1; i >= 0; i-- {
		cur = findCloset(cur, i, target)
		if cur.next[i] != nil && cur.next[i].val == target {
			return true
		}
	}
	return false
}

func (p *Skiplist) Add(num int) {
	level := randomLevel()
	if level > p.level {
		p.level = level
	}
	tmp := newNode(num, level)
	cur := p.head
	for i := p.level - 1; i >= 0; i-- {
		cur = findCloset(cur, i, num)
		if i < level {
			tmp.next[i] = cur.next[i]
			cur.next[i] = tmp
		}
	}
}

func (p *Skiplist) Erase(num int) bool {
	ok := false
	cur := p.head
	for i := p.level - 1; i >= 0; i-- {
		cur = findCloset(cur, i, num)
		if cur.next[i] != nil && cur.next[i].val == num {
			cur.next[i] = cur.next[i].next[i]
			ok = true
		}
	}
	for p.level > 1 && p.head.next[p.level-1] == nil {
		p.level--
	}
	return ok
}

func findCloset(cur *node, level, target int) *node {
	for cur.next[level] != nil && cur.next[level].val < target {
		cur = cur.next[level]
	}
	return cur
}

func randomLevel() int {
	level := 1
	for level < maxLevel && rng.Float64() < p {
		level++
	}
	return level
}

/**
 * Your Skiplist object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Search(target);
 * obj.Add(num);
 * param_3 := obj.Erase(num);
 */

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
