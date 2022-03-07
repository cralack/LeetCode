package mapsumpairs

import "testing"

type TrieNode struct {
	Val      int
	IsEnd    bool
	Children [26]*TrieNode
}
type MapSum struct {
	Root *TrieNode
}

func Constructor() MapSum {
	return MapSum{&TrieNode{}}
}
func (_this *MapSum) prefix(key string) *TrieNode {
	cur := _this.Root
	for _, ch := range key {
		ch -= 'a'
		if cur.Children[ch] == nil {
			cur.Children[ch] = &TrieNode{}
		}
		cur = cur.Children[ch]
	}
	return cur
}
func (_this *MapSum) Insert(key string, val int) {
	node := _this.prefix(key)
	node.Val = val
	node.IsEnd = true
}
func dfs(root *TrieNode) int {
	if root == nil {
		return 0
	}
	sum := root.Val
	for i := 0; i < 26; i++ {
		sum += dfs(root.Children[i])
	}
	return sum
}
func (_this *MapSum) Sum(prefix string) int {
	node := _this.prefix(prefix)
	return dfs(node)
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
func Test_map_sum_pairs(t *testing.T) {
	c1 := []string{"MapSum", "insert", "sum", "insert", "sum"}
	c2 := []string{"", "apple", "ap", "app", "ap"}
	c3 := []int{-1, 3, -1, 2, -1}
	var map_s MapSum
	for idx, c := range c1 {
		switch c {
		case "MapSum":
			map_s = Constructor()
		case "insert":
			map_s.Insert(c2[idx], c3[idx])
		case "sum":
			t.Log(map_s.Sum(c2[idx]))
		}
	}
}
