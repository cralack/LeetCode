package scoreofparentheses

import "testing"

func scoreOfParentheses(s string) int {
	ans, d := 0, 0
	for i, c := range s {
		if c == '(' {
			d++
		} else {
			d--
			if s[i-1] == '(' {
				ans += 1 << d
			}
		}
	}
	return ans
}
func Test_score_of_parentheses(t *testing.T) {
	s := "()"
	t.Log(scoreOfParentheses(s))
	s = "(())"
	t.Log(scoreOfParentheses(s))
	s = "()()"
	t.Log(scoreOfParentheses(s))
	s = "(()(()))"
	t.Log(scoreOfParentheses(s))
}
