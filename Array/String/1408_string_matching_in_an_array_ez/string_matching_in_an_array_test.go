package stringmatchinginanarrayez

import (
	"sort"
	"testing"
)

func stringMatching(words []string) (ans []string) {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isMatch(words[i], words[j]) {
				if len(ans) > 0 && ans[len(ans)-1] == words[i] { //ans 去重
					continue
				}
				ans = append(ans, words[i])
			}
		}
	}
	return
}
func isMatch(a, b string) bool {
	n := len(a)
	for i := 0; i < len(b)-n+1; i++ {
		tmp := b[i : i+n]
		if tmp == a {
			return true
		}
	}
	return false
}
func Test_string_matching_in_an_array(t *testing.T) {
	words := []string{"mass", "as", "hero", "superhero"}
	t.Log(stringMatching(words))
	words = []string{"blue", "green", "bu"}
	t.Log(stringMatching(words))
	words = []string{"leetcoder", "leetcode", "od", "hamlet", "am"}
	t.Log(stringMatching(words))
	words = []string{"leetcode", "et", "code"}
	t.Log(stringMatching(words))
}
