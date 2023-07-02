package oneawaylcci

import (
	"testing"
)

func oneEditAway(first string, second string) bool {
	n, m := len(first), len(second)
	if n < m { // keep len(first)>len(second)
		return oneEditAway(second, first)
	}
	if n-m > 1 {
		return false
	}
	i, j := 0, 0
	cnt := 0
	for i < n && j < m && cnt <= 1 {
		if first[i] == second[j] {
			i++
			j++
		} else {
			if n == m {
				i++
				j++
				cnt++
			} else {
				i++
				cnt++
			}
		}
	}
	return cnt <= 1
}
func Test_one_away_lcci(t *testing.T) {
	first, second := "pale", "ple"
	t.Log(oneEditAway(first, second))
	first, second = "a", "b"
	t.Log(oneEditAway(first, second))
	first, second = "intention", "execution"
	t.Log(oneEditAway(first, second))
	first, second = "pales", "pal"
	t.Log(oneEditAway(first, second))
	first, second = "spartan", "part"
	t.Log(oneEditAway(first, second))
	first, second = "islander", "slander"
	t.Log(oneEditAway(first, second))
}
