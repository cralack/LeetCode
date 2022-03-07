package designaddandsearchwordsdatastructure

import (
	"testing"
)

type WordDictionary struct {
	root *TrieNode
}
type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

func Constructor() WordDictionary {
	return WordDictionary{&TrieNode{}}
}

func (_this *WordDictionary) AddWord(word string) {
	cur := _this.root
	for _, char := range word {
		char -= 'a'
		if cur.children[char] == nil {
			cur.children[char] = &TrieNode{}
		}
		cur = cur.children[char]
	}
	cur.isEnd = true
}

func (_this *WordDictionary) Search(pattern string) bool {
	return dfs(_this.root, pattern, 0)
}
func dfs(node *TrieNode, word string, idx int) bool {
	if idx == len(word) {
		return node.isEnd
	}
	if word[idx] == '.' {
		for i := 0; i < 26; i++ {
			if node.children[i] != nil && dfs(node.children[i], word, idx+1) {
				return true
			}
		}
	} else if node.children[word[idx]-'a'] != nil {
		return dfs(node.children[word[idx]-'a'], word, idx+1)
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
func Test_design_add_and_search_words_data_structure(t *testing.T) {
	c1 := []string{
		"WordDictionary", "addWord", "addWord", "addWord",
		"search", "search", "search", "search"}
	c2 := []string{"", "bad", "dad", "mad",
		"pad", "bad", ".ad", "b.."}
	var dic WordDictionary
	for idx, val := range c1 {
		switch val {
		case "WordDictionary":
			dic = Constructor()
		case "addWord":
			dic.AddWord(c2[idx])
		case "search":
			t.Log(dic.Search(c2[idx]))
		}
	}
}
