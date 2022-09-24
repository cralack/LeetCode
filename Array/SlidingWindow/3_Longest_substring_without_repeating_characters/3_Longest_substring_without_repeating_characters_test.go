package longest_sub_str

import (
	"testing"
)

func lengthOfLongestSubstring(source string) int {
	if len(source) <= 0 {
		return 0
	}
	window := make(map[byte]int)
	left, right := 0, 0
	res := 0
	for right < len(source) {
		c := source[right]
		right++
		window[c]++
		for window[c] > 1 {
			d := source[left]
			left++
			window[d]--
		}
		if res < right-left {
			res = right - left
		}
	}
	return res
}

func TestLonSubStr(t *testing.T) {
	strs := []string{"abcabcbb", "bbbbb", "pwwkew"}
	ans := []int{}
	for _, str := range strs {
		ans = append(ans, lengthOfLongestSubstring(str))
	}
	t.Log(ans)
}
