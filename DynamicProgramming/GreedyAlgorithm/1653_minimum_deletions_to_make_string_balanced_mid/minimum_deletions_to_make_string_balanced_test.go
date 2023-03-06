package minimumdeletions

import (
	"testing"
)

func minimumDeletions(s string) (ans int) {
	//考虑s的最后一个字母
	cntB := 0
	for _, c := range s {
		//无需删除，问题规模缩小，变成使s的前n-1个字母平衡的最少删除次数
		if c == 'b' {
			cntB++
		} else {
			//删除它，则答案为ans+1
			//或保留它，那么前面的所有'b'都要删除
			ans = min(ans+1, cntB)
		}
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_minimum_deletions_to_make_string_balanced(t *testing.T) {
	s := "aababbab"
	t.Log(minimumDeletions(s))
	s = "bbaaaaabb"
	t.Log(minimumDeletions(s))
}
