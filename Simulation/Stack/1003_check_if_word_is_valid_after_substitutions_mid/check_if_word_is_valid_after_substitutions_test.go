package checkifwordisvalidaftersubstitutions

import "testing"

func isValid(s string) bool {
	stack := make([]rune, 0)
	for _, char := range s {
		if char > 'a' {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if char-top != 1 {
				return false
			}
		}
		if char < 'c' {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}

func Test_check_if_word_is_valid_after_substitutions(t *testing.T) {
	tests := []struct {
		s     string
		wanna bool
	}{
		{"aabcbc", true},
		{"abcabcababcc", true},
		{"abccba", false},
	}
	for _, tt := range tests {
		t.Log(isValid(tt.s) == tt.wanna)
	}
}
