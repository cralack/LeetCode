package longeststringchain

import (
	"sort"
	"testing"
)

func longestStrChain(words []string) (ans int) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	dic := make(map[string]int, 0)
	for _, word := range words {
		res := 0
		for i := range word {
			tmp := word[:i] + word[i+1:]
			res = max(res, dic[tmp])
		}
		dic[word] = res + 1
		ans = max(ans, res+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Test_longest_string_chain(t *testing.T) {
	tests := []struct {
		words []string
		wanna int
	}{
		{[]string{"a", "b", "ba", "bca", "bda", "bdca"}, 4},
		{[]string{"xbc", "pcxbcf", "xb", "cxbc", "pcxbc"}, 5},
		{[]string{"abcd", "dbqca"}, 1},
	}
	for _, tt := range tests {
		t.Log(longestStrChain(tt.words) == tt.wanna)
	}
}
