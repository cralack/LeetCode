package maximumscoreaftersplittingastringez

import "testing"

func maxScore(s string) (ans int) {
	sum := 0
	for _, c := range s {
		if c == '1' {
			sum++
		}
	}
	for i := 0; i < len(s)-1; i++ { // 分割成两个 非空 子字符串
		if s[i] == '0' {
			sum++
		} else {
			sum--
		}

		if ans < sum {
			ans = sum
		}
	}
	return ans
}
func Test_maximum_score_after_splitting_a_string(t *testing.T) {
	s := "011101"
	t.Log(maxScore(s))
	s = "00111"
	t.Log(maxScore(s))
	s = "1111"
	t.Log(maxScore(s))
	s = "00"
	t.Log(maxScore(s))
}
