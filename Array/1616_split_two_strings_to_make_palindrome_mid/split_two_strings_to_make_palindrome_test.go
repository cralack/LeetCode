package splittwostringstomakepalindrome

import "testing"

func checkPalindromeFormation(a string, b string) bool {
	return check(a, b) || check(b, a)
}

// 如果从 s[i] 到 s[j] 是回文串则返回 true，否则返回 false
func isPalindrome(s string, i, j int) bool {
	for i < j && s[i] == s[j] {
		i++
		j--
	}
	return i >= j
}

// 如果 a_prefix + b_suffix 可以构成回文串则返回 true，否则返回 false
func check(a, b string) bool {
	i, j := 0, len(a)-1         // 相向双指针
	for i < j && a[i] == b[j] { // 尽量匹配前后缀
		i++
		j--
	}
	return isPalindrome(a, i, j) || isPalindrome(b, i, j)
}
func Test_split_two_strings_to_make_palindrome(t *testing.T) {
	tests := []struct{ a, b string }{
		{a: "x", b: "y"},
		{a: "abdef", b: "fecab"},
		{a: "ulacfd", b: "jizalu"},
	}
	for _, tt := range tests {
		t.Log(checkPalindromeFormation(tt.a, tt.b))
	}
}
