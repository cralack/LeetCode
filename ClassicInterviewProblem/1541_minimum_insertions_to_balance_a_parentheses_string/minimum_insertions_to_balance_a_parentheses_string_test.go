package minimuminsertionstobalanceaparenthesesstring

import "testing"

func minInsertions(s string) int {
	needLeft, needRight := 0, 0

	for _, c := range s {
		// 一个左括号对应两个右括号
		if c == '(' {
			needRight += 2
			if needRight%2 == 1 {
				needLeft++
				needRight--
			}
		}
		if c == ')' {
			needRight--
			if needRight == -1 { // 右括号多了
				needLeft++    // 需要插入一个左括号
				needRight = 1 // 同时，对右括号的需求变为 1
			}
		}
	}

	return needLeft + needRight
}
func Test_minimum_insertions_to_balance_a_parentheses_string(t *testing.T) {
	t.Log(minInsertions("(()))"))
	t.Log(minInsertions("())"))
	t.Log(minInsertions("))())("))
	t.Log(minInsertions("(((((("))
	t.Log(minInsertions(")))))))"))
}
