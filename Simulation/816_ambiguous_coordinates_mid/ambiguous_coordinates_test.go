package ambiguouscoordinates

import "testing"

func ambiguousCoordinates(s string) (ans []string) {
	n := len(s)
	s = s[1 : n-1]
	n -= 2
	for idx := 1; idx < n; idx++ {
		left := split(s[:idx])
		right := split(s[idx:])
		for _, lstr := range left {
			for _, rstr := range right {
				ans = append(ans, "("+lstr+", "+rstr+")") // wdnmd逗号后面还有个空格呢?
			}
		}
	}
	return ans
}

func split(str string) (res []string) {
	n := len(str)
	for k := 1; k <= n; k++ {
		l, r := str[0:k], str[k:n]
		// 左半部分不能以 00 开头，除非左半部分只有 00
		ok := (l == "0" || l[0] != '0') &&
			// 右半部分不能以 00 结尾。
			(r == "" || r[len(r)-1] != '0')
		if ok {
			op := ""
			if k < n {
				op = "."
			}
			res = append(res, l+op+r)
		}
	}
	return
}

func Test_ambiguous_coordinates(t *testing.T) {
	s := "(123)"
	t.Log(ambiguousCoordinates(s))
	s = "(00011)"
	t.Log(ambiguousCoordinates(s))
	s = "(0123)"
	t.Log(ambiguousCoordinates(s))
	s = "(100)"
	t.Log(ambiguousCoordinates(s))
}
