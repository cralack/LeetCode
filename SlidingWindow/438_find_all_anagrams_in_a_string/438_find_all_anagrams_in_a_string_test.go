package findallanagramsinastring

import (
	"testing"
)

func findAnagrams(source, target string) (res []int) {
	if len(source) < len(target) {
		return nil
	}
	need, window := [26]int{}, [26]int{}
	for _, char := range target {
		need[char-'a']++
	}
	left, right := 0, 0
	valid := 0

	for right < len(source) {
		c := source[right]
		right++
		if need[c-'a'] > 0 {
			window[c-'a']++
			if window[c-'a'] <= need[c-'a'] {
				valid++
			}
		}

		for right-left >= len(target) {
			if valid == len(target) {
				res = append(res, left)
			}
			d := source[left]
			left++
			if need[d-'a'] > 0 {
				if window[d-'a'] <= need[d-'a'] {
					valid--
				}
				window[d-'a']--
			}
		}
	}
	return res
}
func Test_find_all_anagrams_in_a_string(t *testing.T) {
	s := []string{"cbaebabacd", "abab", "baa"}
	p := []string{"abc", "ab", "aa"}
	for i, str := range s {
		res := findAnagrams(str, p[i])
		t.Log(res)
	}
}
