package take_k_of_each_character_from_left_and_right_mid

import (
	"testing"
)

func takeCharacters(s string, k int) int {
	cnt := make([]int, 3)
	for _, char := range s {
		cnt[char-'a']++
	}
	if cnt[0] < k || cnt[1] < k || cnt[2] < k {
		return -1 // 字母个数不足 k
	}

	mx, left := 0, 0
	for right, char := range s {
		char -= 'a'
		cnt[char]--
		for cnt[char] < k {
			cnt[s[left]-'a']++
			left++
		}
		mx = max(mx, right-left+1)
	}
	return len(s) - mx
}

func Test_take_k_of_each_character_from_left_and_right(t *testing.T) {
	tests := []struct {
		s string
		k int
	}{
		{s: "aabaaaacaabc", k: 2},
		{s: "a", k: 1},
	}
	for _, test := range tests {
		t.Log(takeCharacters(test.s, test.k))
	}
}
