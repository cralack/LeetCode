package evaluatethebracketpairsofastring

import (
	"testing"
)

func evaluate(s string, knowledge [][]string) string {
	dic := make(map[string]string, len(knowledge))
	n := len(s)
	ans := make([]byte, 0, n)
	for _, know := range knowledge {
		dic[know[0]] = know[1]
	}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			j := i + 1
			for j < n && s[j] != ')' {
				j++
			}

			tar, has := dic[s[i+1:j]]
			if has {
				ans = append(ans, []byte(tar)...)
			} else {
				ans = append(ans, '?')
			}
			i = j
		} else {
			ans = append(ans, s[i])
		}
	}
	return string(ans)
}

func Test_evaluate_the_bracket_pairs_of_a_string(t *testing.T) {
	s := "(name)is(age)yearsold"
	knowledge := [][]string{{"name", "bob"}, {"age", "two"}}
	t.Log(evaluate(s, knowledge))
	s = "hi(name)"
	knowledge = [][]string{{"a", "b"}}
	t.Log(evaluate(s, knowledge))
	s = "(a)(a)(a)aaa"
	knowledge = [][]string{{"a", "yes"}}
	t.Log(evaluate(s, knowledge))
}
