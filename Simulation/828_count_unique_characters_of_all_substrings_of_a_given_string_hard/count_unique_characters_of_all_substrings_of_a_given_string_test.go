package uniqueLetterString

import "testing"

func uniqueLetterString(s string) (ans int) {
	str := []byte(s)
	n := len(str)
	left, right := make([]int, n), make([]int, n)
	cnts := make([]int, 5)

	for i := range cnts {
		cnts[i] = -1
	}
	// str[i] 能够作为子数组唯一字符时的最远左边界
	for i := 0; i < n; i++ {
		u := int(str[i] - 'A')
		left[i] = cnts[u]
		cnts[u] = i
	}

	for i := range cnts {
		cnts[i] = n
	}
	// str[i] 能够作为子数组唯一字符时的最远右边界
	for i := n - 1; i >= 0; i-- {
		u := int(str[i] - 'A')
		right[i] = cnts[u]
		cnts[u] = i
	}

	for i := 0; i < n; i++ {
		ans += (i - left[i]) * (right[i] - i)
	}
	return
}
func Test_uniqueLetterString(t *testing.T) {
	s := "ABC"
	t.Log(uniqueLetterString(s))
	s = "ABA"
	t.Log(uniqueLetterString(s))
	// s = "LEETCODE"
	// t.Log(uniqueLetterString(s))
}
