package longestuncommonsubsequenceiimid

import (
	"testing"
)

func findLUSlength(strs []string) int {
	maxLen := -1
	for i, str1 := range strs {
		flag := false
		for j, str2 := range strs {
			if i != j && !flag {
				flag = isSubsequence(str1, str2)
			}
		}
		//str1不是其他任意字符串的子序列
		if !flag && len(str1) > maxLen {
			maxLen = len(str1)
		}
	}
	return maxLen
}

func isSubsequence(pat, txt string) bool {
	if len(pat) > len(txt) {
		return false
	}
	j := 0
	for i := 0; i < len(txt); i++ {
		if j >= len(pat) {
			return true
		}
		if pat[j] == txt[i] {
			j++
		}
	}
	return j >= len(pat)
}
func Test_longest_uncommon_subsequence_ii(t *testing.T) {
	strs := []string{"aba", "cdc", "eae"}
	t.Log(findLUSlength(strs))
	strs = []string{"aaa", "aaa", "aa"}
	t.Log(findLUSlength(strs))
}
