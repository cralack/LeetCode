package minimumwindowsubstring

import (
	"testing"
)

func checkInclusion(target, source string) bool {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, c := range target {
		need[byte(c)]++
	}
	left, right := 0, 0
	valid, size := 0, len(source)
	for right < size {
		c := source[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for right-left >= len(target) {
			if valid == len(need) {
				return true
			}
			d := source[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}
func Test_permutation_in_string(t *testing.T) {
	s1 := []string{"ab", "ab"}
	s2 := []string{"eidbaooo", "eidboaoo"}
	for i, val := range s2 {
		ok := checkInclusion(s1[i], val)
		t.Log(ok)
	}
}
