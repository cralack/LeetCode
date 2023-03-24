package streamofcharacters

import (
	"strings"
	"testing"
)

type Trie struct {
	end   bool
	child [26]*Trie
}
type StreamChecker struct {
	sb   *strings.Builder
	root *Trie
}

func Constructor(words []string) StreamChecker {
	r := &Trie{}
	for _, word := range words {
		r.insert(word)
	}
	return StreamChecker{
		&strings.Builder{},
		r,
	}
}

func (this *Trie) insert(word string) {
	cur := this
	//倒序建树
	for i := len(word) - 1; i >= 0; i-- {
		char := word[i] - 'a'
		if cur.child[char] == nil {
			cur.child[char] = &Trie{}
		}
		cur = cur.child[char]
	}
	cur.end = true
}

func (this *Trie) search(word string) bool {
	cur := this
	for i, j := len(word)-1, 0; i >= 0 && j < 201; i, j = i-1, j+1 {
		char := word[i] - 'a'
		if cur.child[char] == nil {
			return false
		}
		cur = cur.child[char]
		if cur.end {
			return true
		}
	}
	return false
}

func (this *StreamChecker) Query(letter byte) bool {
	this.sb.WriteByte(letter)
	return this.root.search(this.sb.String())
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
func Test_stream_of_characters(t *testing.T) {
	tests := []struct {
		words   []string
		queries []byte
	}{{[]string{"cd", "f", "kl"},
		[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l'}},
	}
	for _, tt := range tests {
		sol := Constructor(tt.words)
		for _, char := range tt.queries {
			t.Log(sol.Query(char))
		}
	}
}
