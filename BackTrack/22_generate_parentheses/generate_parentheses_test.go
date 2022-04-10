package generateparentheses

import "testing"

func generateParenthesis(n int) (res []string) {
	if n == 0 {
		return res
	}
	path := []byte{}
	var backtrack func(left, right int)
	backtrack = func(left, right int) {
		// 若左括号剩下的多，说明不合法
		if left > right {
			return
		} // 数量小于 0 肯定是不合法的
		if left < 0 || right < 0 {
			return
		} // 当所有括号都恰好用完时，得到一个合法的括号组合
		if left == 0 && right == 0 {
			res = append(res, string(path))
			return
		}
		// 尝试放一个左括号
		path = append(path, '(')
		backtrack(left-1, right)
		path = path[:len(path)-1]

		// 尝试放一个右括号
		path = append(path, ')')
		backtrack(left, right-1)
		path = path[:len(path)-1]
	}
	backtrack(n, n)
	return res
}
func Test_generate_parentheses(t *testing.T) {
	n := 3
	t.Log(generateParenthesis(n))
}
