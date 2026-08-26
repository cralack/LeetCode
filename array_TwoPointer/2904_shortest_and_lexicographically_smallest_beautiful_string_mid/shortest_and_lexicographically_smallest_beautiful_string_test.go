package shortest_and_lexicographically_smallest_beautiful_string_mid

import (
	"strings"
	"testing"
)

func shortestBeautifulSubstring(s string, k int) string {
	if strings.Count(s, "1") < k {
		return ""
	}
	ans := s
	cnt1 := 0
	left := 0
	for right, b := range s {
		cnt1 += int(b & 1)
		for cnt1 > k || s[left] == '0' {
			cnt1 -= int(s[left] & 1)
			left++
		}
		if cnt1 == k {
			t := s[left : right+1]
			if len(t) < len(ans) || len(t) == len(ans) && t < ans {
				ans = t
			}
		}
	}
	return ans
}

func Test_shortest_and_lexicographically_smallest_beautiful_string(t *testing.T) {
	tests := []struct {
		s string
		k int
	}{
		{s: "100011001", k: 3},
		{s: "1011", k: 2},
		{s: "000", k: 1},
	}
	for _, test := range tests {
		t.Log(shortestBeautifulSubstring(test.s, test.k))
	}
}
