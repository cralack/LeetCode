package comparestringsbyfrequency

import (
	"sort"
	"testing"
)

func numSmallerByFrequency(queries []string, words []string) (ans []int) {
	f := func(str string) int {
		cnt := [26]int{}
		for _, char := range str {
			cnt[char-'a']++
		}
		for _, c := range cnt {
			if c > 0 {
				return c
			}
		}
		return 0
	}
	n := len(words)
	dic := make([]int, n)
	for i, word := range words {
		dic[i] = f(word)
	}
	sort.Ints(dic)
	ans = make([]int, len(queries))
	for i, que := range queries {
		tar := f(que)
		idx := sort.Search(n, func(mid int) bool {
			return dic[mid] > tar
		})
		ans[i] = n - idx
	}
	return
}

func Test_compare_strings_by_frequency_of_the_smallest_character(t *testing.T) {
	tests := []struct {
		queries []string
		words   []string
	}{
		{queries: []string{"cbd"}, words: []string{"zaaaz"}},
		{queries: []string{"bbb", "cc"}, words: []string{"a", "aa", "aaa", "aaaa"}},
	}
	for _, tt := range tests {
		t.Log(numSmallerByFrequency(tt.queries, tt.words))
	}
}
