package minimumaddtomakeparenthesesvalid

import "testing"

func minAddToMakeValid(s string) int {
	needLeft, needRight := 0, 0 //对左右括号的需求量
	for _, c := range s {
		if c == '(' { // 对右括号的需求 + 1
			needRight++
		}
		if c == ')' { // 对右括号的需求 - 1
			needRight--
			if needRight == -1 {
				needRight = 0
				needLeft++ // 需插入一个左括号
			}
		}
	}
	return needLeft + needRight
}
func Test_minimum_add_to_make_parentheses_valid(t *testing.T) {
	t.Log(minAddToMakeValid("())"))
	t.Log(minAddToMakeValid("((("))
	t.Log(minAddToMakeValid("()))(("))
}
