package differentwaystoaddparentheses

import (
	"strconv"
	"testing"
)

var memo = make(map[string][]int, 0)

func diffWaysToCompute(expression string) (res []int) {
	if ans, ok := memo[expression]; ok {
		return ans
	}
	for i := 0; i < len(expression); i++ {
		c := expression[i]
		// 扫描算式 expression 中的运算符
		if c == '+' || c == '-' || c == '*' {
			// 以运算符为中心，分割成两个字符串，分别递归计算
			left := diffWaysToCompute(expression[0:i])
			right := diffWaysToCompute(expression[i+1:])
			// 通过子问题的结果，合成原问题的结果
			for _, a := range left {
				for _, b := range right {
					switch c {
					case '+':
						res = append(res, a+b)
					case '-':
						res = append(res, a-b)
					case '*':
						res = append(res, a*b)
					}
				}
			}
		}
	}
	// base case
	// 如果 res 为空，说明算式是一个数字，没有运算符
	if len(res) == 0 {
		n, err := strconv.Atoi(expression)
		if err == nil {
			res = append(res, n)
		}
	}
	memo[expression] = res
	return
}
func Test_different_ways_to_add_parentheses(t *testing.T) {
	expression := "2-1-1"
	t.Log(diffWaysToCompute(expression))
	expression = "2*3-4*5"
	t.Log(diffWaysToCompute(expression))
}
func Benchmark_different_ways_to_add_parentheses(b *testing.B) {
	expression := "2*3-4*5"
	b.Run("with_memo", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			diffWaysToCompute(expression)
		}
		b.StopTimer()
	})
}
