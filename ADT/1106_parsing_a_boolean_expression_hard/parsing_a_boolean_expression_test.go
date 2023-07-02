package parsingabooleanexpression

import "testing"

func parseBoolExpr(expression string) bool {
	numStk := &runeStack{}
	opStk := &runeStack{}
	for _, char := range expression {
		switch char {
		case 't', 'f':
			numStk.Push(char)
		case '|', '&', '!':
			opStk.Push(char)
		case '(':
			numStk.Push('-')
		case ')':
			op := opStk.Pop()
			cur := ' '
			// 连续计算一个括号内
			for !numStk.IsEmpty() && numStk.Peak() != '-' {
				top := numStk.Pop()
				if cur == ' ' {
					cur = top
				} else {
					cur = calc(top, cur, op)
				}
			}
			if op == '!' && cur == 't' {
				cur = 'f'
			} else if op == '!' {
				cur = 't'
			}
			numStk.Pop() // 取出'-'
			numStk.Push(cur)
		}
	}
	return numStk.Peak() == 't'
}

// 计算括号内两个操作数结果
func calc(a, b, op rune) rune {
	x := a == 't'
	y := b == 't'
	var ans bool

	if op == '|' {
		ans = x || y
	} else {
		ans = x && y
	}
	if ans {
		return 't'
	}
	return 'f'
}

// stack封装
type runeStack []rune

func (s runeStack) Peak() rune     { return s[len(s)-1] }
func (s *runeStack) Len() int      { return len(*s) }
func (s *runeStack) IsEmpty() bool { return len(*s) == 0 }
func (s *runeStack) Push(tar rune) { *s = append(*s, tar) }
func (s *runeStack) Pop() rune {
	old := *s
	tar := old[len(old)-1]
	*s = old[:len(old)-1]
	return tar
}

func Test_parsing_a_boolean_expression(t *testing.T) {
	expression := "!(f)"
	t.Log(parseBoolExpr(expression))
	expression = "|(f,t)"
	t.Log(parseBoolExpr(expression))
	expression = "&(t,f)"
	t.Log(parseBoolExpr(expression))
	expression = "|(&(t,f,t),!(t))"
	t.Log(parseBoolExpr(expression))
}
