package checkifonestringswapcanmakestringsequal

import "testing"

func areAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	n := len(s1)
	cnt := 0
	a, b := -1, -1
	for i := 0; i < n; i++ {
		if s1[i] != s2[i] {
			if a < 0 {
				a = i
			} else if b < 0 {
				b = i
			}
			cnt++
		}
	}
	return cnt == 2 && s1[a] == s2[b] && s2[a] == s1[b]
}
func Test_almostEqual(t *testing.T) {
	s1 := "bank"
	s2 := "kanb"
	t.Log(areAlmostEqual(s1, s2))
	s1 = "attack"
	s2 = "defend"
	t.Log(areAlmostEqual(s1, s2))
	s1 = "kelb"
	s2 = "kelb"
	t.Log(areAlmostEqual(s1, s2))
	s1 = "abcd"
	s2 = "dcba"
	t.Log(areAlmostEqual(s1, s2))
	s1 = "aa"
	s2 = "ac"
	t.Log(areAlmostEqual(s1, s2))
	s1 = "caa"
	s2 = "aaz"
	t.Log(areAlmostEqual(s1, s2))
}
