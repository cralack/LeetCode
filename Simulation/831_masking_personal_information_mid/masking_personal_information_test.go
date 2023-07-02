package maskingpersonalinformation

import (
	"strings"
	"testing"
)

func maskPII(tar string) string {
	if idx := strings.Index(tar, "@"); idx != -1 {
		// s is mail add
		ans := make([]byte, 0, 2*len(tar))
		ans = append(ans, toLower(tar[0]))
		ans = append(ans, []byte("*****")...)
		for i := idx - 1; i < len(tar); i++ {
			ans = append(ans, toLower(tar[i]))
		}
		return string(ans)
	} else {
		// s is phone num
		num := make([]byte, 0, len(tar))
		for _, c := range tar {
			if '0' <= c && c <= '9' {
				num = append(num, byte(c))
			}
		}
		suf := "***-***-" + string(num[len(num)-4:])
		if len(num) == 10 {
			return suf
		}
		return "+" + strings.Repeat("*", len(num)-10) + "-" + suf
	}
}
func toLower(char byte) byte {
	if 'A' <= char && char <= 'Z' {
		char = char - 'A' + 'a'
	}
	return char
}
func Test_masking_personal_information(t *testing.T) {
	strs := []string{"LeetCode@LeetCode.com", "AB@qq.com",
		"1(234)567-890", "86-(10)12345678"}
	for _, s := range strs {
		t.Log(maskPII(s))
	}
}
