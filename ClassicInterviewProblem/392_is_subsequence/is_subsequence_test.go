package issubsequence

import (
	"testing"
)

// 因为会有make操作，所以实际效率不如理论效率
//
//	func isSubsequence(pat, txt string) bool {
//		index := make(map[byte][]int, 0)
//		for i := 0; i < len(txt); i++ {
//			c := txt[i]
//			index[c] = append(index[c], i)
//		}
//		for i := 0; i < len(pat); i++ {
//			arr, ok := index[pat[i]]
//			if !ok {
//				return false
//			}
//			lo, hi := 0, len(arr)
//			for lo < hi {
//				mi := lo + (hi-lo)>>1
//				if arr[mi] < i {
//					lo = mi + 1
//				} else {
//					hi = mi
//				}
//			}
//			if lo == len(index[pat[i]]) {
//				return false
//			}
//		}
//		return true
//	}
func isSubsequence(pat, txt string) bool {
	i, j := 0, 0
	for i < len(pat) && j < len(txt) {
		if pat[i] == txt[j] {
			i++
		}
		j++
	}
	return i == len(pat)
}
func Test_is_subsequence(t *testing.T) {
	s, p := "abc", "ahbgdc"
	t.Log(isSubsequence(s, p))
	s, p = "axc", "ahbgdc"
	t.Log(isSubsequence(s, p))
}
