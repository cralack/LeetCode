package differentwaystoaddparenthesesmid

import (
	"strconv"
	"testing"
)

func diffWaysToCompute(expression string) []int {
	var dfs func(string) []int

	dfs = func(str string) (ans []int) {
		if val, err := strconv.Atoi(str); err == nil {
			// 如果是数字，直接返回
			return []int{val}
		} else {
			for index, char := range str {
				if char == '+' || char == '-' || char == '*' {
					// 遇到运算符，则左右两边分治
					left := dfs(str[:index])
					right := dfs(str[index+1:])

					for _, leftNum := range left {
						for _, rightNum := range right {
							var res int
							switch char {
							case '+':
								res = leftNum + rightNum
							case '-':
								res = leftNum - rightNum
							case '*':
								res = leftNum * rightNum
							}
							ans = append(ans, res)
						}
					}
				}
			}
		}
		return
	}

	return dfs(expression)
}
func Test_different_ways_to_add_parentheses(t *testing.T) {
	expression := "2-1-1"
	t.Log(diffWaysToCompute(expression))
	expression = "2*3-4*5"
	t.Log(diffWaysToCompute(expression))
}
