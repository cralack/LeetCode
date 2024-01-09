package extra_characters_in_a_string_mid

import (
	"testing"
)

type Trie struct {
	children [26]*Trie
	end      bool
}

func minExtraChar(s string, dictionary []string) int {
	root := &Trie{}
	for _, word := range dictionary {
		cur := root
		for k := len(word) - 1; k >= 0; k-- {
			i := int(word[k] - 'a')
			if cur.children[i] == nil {
				cur.children[i] = &Trie{}
			}
			cur = cur.children[i]
		}
		cur.end = true
	}

	n := len(s)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = f[i-1] + 1
		cur := root
		for j := i - 1; j >= 0; j-- {
			cur = cur.children[int(s[j]-'a')]
			if cur == nil {
				break
			}
			if cur.end && f[j] < f[i] {
				f[i] = f[j]
			}
		}
	}
	return f[n]
}

func Test_extra_characters_in_a_string(t *testing.T) {
	tests := []struct {
		s          string
		dictionary []string
	}{
		{s: "leetscode", dictionary: []string{"leet", "code", "leetcode"}},
		{s: "sayhelloworld", dictionary: []string{"hello", "world"}},
	}
	for _, tt := range tests {
		t.Log(minExtraChar(tt.s, tt.dictionary))
	}
}
