package longest_common_suffix_queries_hard

import (
	"math"
	"testing"
)

func stringIndices(wordsContainer, wordsQuery []string) []int {
	type node struct {
		son       [26]*node
		minLen    int // 子树中的最短字符串的长度
		bestIndex int // 子树中的最短字符串的下标
	}
	root := &node{minLen: math.MaxInt}

	for i, s := range wordsContainer {
		l := len(s)
		if l < root.minLen {
			root.minLen = l
			root.bestIndex = i
		}

		// 把 reverse(s) 插入字典树
		cur := root
		for j := l - 1; j >= 0; j-- {
			b := s[j] - 'a'
			if cur.son[b] == nil {
				cur.son[b] = &node{minLen: math.MaxInt}
			}
			cur = cur.son[b]
			// 维护 cur 子树中的最短字符串的长度及其下标
			// 由于我们是按照 i 从小到大的顺序遍历，字符串长度相同时不更新 bestIndex
			if l < cur.minLen {
				cur.minLen = l
				cur.bestIndex = i
			}
		}
	}

	ans := make([]int, len(wordsQuery))
	for i, s := range wordsQuery {
		cur := root
		for j := len(s) - 1; j >= 0 && cur.son[s[j]-'a'] != nil; j-- {
			cur = cur.son[s[j]-'a']
		}
		// 退出循环时，cur 即最长公共前缀（的对应节点），cur.bestIndex 是前缀为 cur 的最短字符串的下标
		ans[i] = cur.bestIndex
	}
	return ans
}

func Test_longest_common_suffix_queries(t *testing.T) {
	tests := []struct {
		wordsContainer []string
		wordsQuery     []string
	}{
		{wordsContainer: []string{"abcd", "bcd", "xbcd"}, wordsQuery: []string{"cd", "bcd", "xyz"}},
		{wordsContainer: []string{"abcdefgh", "poiuygh", "ghghgh"}, wordsQuery: []string{"gh", "acbfgh", "acbfegh"}},
	}
	for _, tt := range tests {
		t.Log(stringIndices(tt.wordsContainer, tt.wordsQuery))
	}
}
