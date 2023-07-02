package removeoutermostparentheses

import "testing"

func removeOuterParentheses(s string) string {
	var ans, stack []byte
	for _, char := range s {
		if char == ')' {
			stack = stack[:len(stack)-1] // pop
		}
		if len(stack) > 0 {
			ans = append(ans, byte(char))
		}
		if char == '(' {
			stack = append(stack, byte(char))
		}
	}
	return string(ans)
}
func Test_remove_outermost_parentheses(t *testing.T) {
	s := "(()())(())"
	t.Log(removeOuterParentheses(s))
	s = "(()())(())(()(()))"
	t.Log(removeOuterParentheses(s))
	s = "()()"
	t.Log(removeOuterParentheses(s))
}
