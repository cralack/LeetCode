package find_longest_special_substring_that_occurs_thrice_i_mid

import (
	"slices"
	"testing"
)

func maximumLength(s string) int {
	groups := [26][]int{}
	cnt := 0
	for i := range s {
		cnt++
		if i+1 == len(s) || s[i] != s[i+1] {
			groups[s[i]-'a'] = append(groups[s[i]-'a'], cnt) // 统计连续字符长度
			cnt = 0
		}
	}

	ans := 0
	for _, a := range groups {
		if len(a) == 0 {
			continue
		}
		slices.SortFunc(a, func(a, b int) int { return b - a })
		a = append(a, 0, 0) // 假设还有两个空串
		ans = max(ans, a[0]-2, min(a[0]-1, a[1]), a[2])
	}

	if ans == 0 {
		return -1
	}
	return ans
}

func Test_find_longest_special_substring_that_occurs_thrice_i(t *testing.T) {
	tests := []struct{ s string }{
		{s: "aaaa"},
		{s: "abcdef"},
		{s: "abcaba"},
	}
	for _, test := range tests {
		t.Log(maximumLength(test.s))
	}
}
