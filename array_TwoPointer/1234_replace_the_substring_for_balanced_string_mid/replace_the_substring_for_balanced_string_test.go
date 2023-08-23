package replacethesubstringforbalancedstring

import "testing"

func balancedString(s string) int {
	cnt, m := make(map[rune]int, 0), len(s)/4
	for _, c := range s {
		cnt[c]++
	}
	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0
	}
	ans, left := len(s), 0
	for right, c := range s {
		cnt[c]--
		for cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m {
			ans = min(ans, right-left+1)
			cnt[rune(s[left])]++
			left++
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func Test_replace_the_substring_for_balanced_string(t *testing.T) {
	s := "QWER"
	t.Log(balancedString(s))
	s = "QQWE"
	t.Log(balancedString(s))
	s = "QQQW"
	t.Log(balancedString(s))
	s = "QQQQ"
	t.Log(balancedString(s))
}
