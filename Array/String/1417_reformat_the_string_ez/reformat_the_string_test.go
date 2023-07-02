package reformatthestringez

import (
	"testing"
)

func reformat(s string) string {
	size := len(s)
	cntDig, cntStr := 0, 0
	for _, c := range s {
		if isDigit(byte(c)) {
			cntDig++
		}
	}
	cntStr = size - cntDig
	if abs(cntDig-cntStr) > 1 {
		return ""
	}

	flag := cntDig > cntStr // 判断偶数位放字母还是数字
	ans := []byte(s)
	for i, j := 0, 1; i < len(ans); i += 2 {
		if isDigit(ans[i]) != flag { // 偶数位放置错误
			for isDigit(ans[j]) != flag { // 奇数位放置错误
				j += 2
			}
			ans[i], ans[j] = ans[j], ans[i]
		}
	}
	return string(ans)
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
func Test_reformat_the_string(t *testing.T) {
	s := "a0b1c2"
	t.Log(reformat(s))
	s = "leetcode"
	t.Log(reformat(s))
	s = "1229857369"
	t.Log(reformat(s))
	s = "covid2019"
	t.Log(reformat(s))
	s = "ab123"
	t.Log(reformat(s))
}
