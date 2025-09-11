package sort_vowels_in_a_string_mid

import (
	"slices"
	"testing"
)

const vowelMask = 0x208222

func sortVowels(s string) string {
	ans := []byte(s)
	var tmp []byte
	var idx []int
	for i, ch := range []byte(s) {
		if vowelMask>>(ch&31)&1 > 0 {
			idx = append(idx, i)
			tmp = append(tmp, ch)
		}
	}
	if len(idx) == 0 {
		return s
	}
	slices.Sort(tmp)
	for i, j := range idx {
		ans[j] = tmp[i]
	}
	return string(ans)
}

func Test_sort_vowels_in_a_string(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "lEetcOde", want: "lEOtcede"},
		{s: "lYmpH", want: "lYmpH"},
	}
	for _, test := range tests {
		t.Log(sortVowels(test.s) == test.want)
	}
}
