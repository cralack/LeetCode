package basiccalculator

import (
	"testing"
)

func calculate(s string) int {
	var i = 0
	isDigit := func(c byte) bool {
		return '0' <= c && c <= '9'
	}
	var helper func(str string) int
	helper = func(str string) int {
		stack := []int{}
		num := 0          // 记录算式中的数字
		sign := byte('+') // 记录 num 前的符号，初始化为 +
		for ; i < len(s); i++ {
			idx := i
			c := s[idx]
			// char := string(c)
			// fmt.Println(char)
			if isDigit(c) { // 如果是数字，连续读取到 num
				num = 10*num + int(c-'0')
			}
			if c == '(' {
				i++
				num = helper(s)
			}
			// 如果不是数字，就是遇到了下一个符号，
			// 之前的数字和符号就要存进栈中
			if !isDigit(c) && c != ' ' || i == len(s)-1 {
				switch sign {
				case '+':
					stack = append(stack, num)
				case '-':
					stack = append(stack, -num)
				case '*':
					stack[len(stack)-1] *= num
				case '/':
					stack[len(stack)-1] /= num
				}
				// 更新符号为当前符号，数字清零
				sign = c
				num = 0
				if c == ')' {
					break
				}
			}
		}
		res := 0
		for _, n := range stack {
			res += n
		}
		return res
	}
	return helper(s)
}
func Test_basic_calculator(t *testing.T) {
	s := "1 + 1"
	t.Log(calculate(s))
	s = " 2-1 + 2 "
	t.Log(calculate(s))
	s = "(1+(4+5+2)-3)+(6+8)"
	t.Log(calculate(s))
	s = "3 * (4 - 5 / 2 ) - 6"
	t.Log(calculate(s))
}
