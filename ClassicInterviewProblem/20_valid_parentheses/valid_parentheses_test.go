package validparentheses

import "testing"

func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			var pair rune
			switch c {
			case ')':
				pair = '('
			case ']':
				pair = '['
			case '}':
				pair = '{'
			}
			if len(stack) != 0 && stack[len(stack)-1] == pair {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
func Test_valid_parentheses(t *testing.T) {
	t.Log(isValid("()"))
	t.Log(isValid("()[]{}"))
	t.Log(isValid("(]"))
	t.Log(isValid("([)]"))
	t.Log(isValid("{[]}"))
	t.Log(isValid("]"))
}
