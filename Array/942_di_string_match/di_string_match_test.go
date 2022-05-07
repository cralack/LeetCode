package distringmatch

import (
	"testing"
)

func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	left, right := 0, n
	for i, c := range s {
		if c == 'I' {
			res[i] = left
			left++
		} else {
			res[i] = right
			right--
		}
		i++
	}
	res[n] = left
	return res
}
func Test_di_string_match(t *testing.T) {
	s := "IDID"
	t.Log(diStringMatch(s))
	s = "III"
	t.Log(diStringMatch(s))
	s = "DDI"
	t.Log(diStringMatch(s))
}
