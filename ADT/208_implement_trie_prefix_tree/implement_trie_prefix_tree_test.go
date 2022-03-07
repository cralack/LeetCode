package implementtrieprefixtree

import "testing"

type Trie struct {
	IsEnd    bool
	Children [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

func (_this *Trie) Insert(word string) {
	cur := _this
	for _, ch := range word {
		ch -= 'a'
		if cur.Children[ch] == nil {
			cur.Children[ch] = &Trie{}
		}
		cur = cur.Children[ch]
	}
	cur.IsEnd = true
}
func (_this *Trie) SearchPrefix(word string) *Trie {
	cur := _this
	for _, ch := range word {
		ch -= 'a'
		if cur.Children[ch] == nil {
			return nil
		}
		cur = cur.Children[ch]
	}
	return cur
}
func (_this *Trie) Search(word string) bool {
	node := _this.SearchPrefix(word)
	return node != nil && node.IsEnd
}

func (_this *Trie) StartsWith(prefix string) bool {
	return _this.SearchPrefix(prefix) != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func Test_implement_trie_prefix_tree(t *testing.T) {
	c1 := []string{"Trie", "insert", "search", "search",
		"startsWith", "insert", "search"}
	c2 := []string{"", "apple", "apple", "app", "app", "app", "app"}
	// c1 := []string{"Trie", "insert", "search", "startsWith"}
	// c2 := []string{"", "a", "a", "a"}
	var root Trie
	for idx, val := range c1 {
		switch val {
		case "Trie":
			root = Constructor()
		case "insert":
			root.Insert(c2[idx])
		case "search":
			t.Log(root.Search(c2[idx]))
		case "startsWith":
			t.Log(root.StartsWith(c2[idx]))
		}
	}

}
