package check_if_strings_can_be_made_equal_with_operations_ii_mid

import (
	"testing"
)

func checkStrings(s1 string, s2 string) bool {
	cnt1 := [2][26]int{}
	cnt2 := [2][26]int{}
	for i := range s1 {
		cnt1[i%2][s1[i]-'a']++
		cnt2[i%2][s2[i]-'a']++
	}
	return cnt1 == cnt2
}

func Test_check_if_strings_can_be_made_equal_with_operations_ii(t *testing.T) {
	tests := []struct {
		s1, s2 string
	}{
		{s1: "abcdba", s2: "cabdab"},
		{s1: "abe", s2: "bea"},
		{s1: "kvwdssgl", s2: "wskxsdgv"},
		{s1: "z", s2: "u"},
	}
	for _, tt := range tests {
		t.Log(checkStrings(tt.s1, tt.s2))
	}
}
