package process_string_with_special_operations_i_mid

import (
	"testing"
)

func processStr(s string) string {
	ans := make([]byte, 0, len(s))
	for _, ch := range []byte(s) {
		switch ch {
		case '*':
			ans = ans[:max(len(ans)-1, 0)]
		case '#':
			ans = append(ans, ans...)
		case '%':
			for i, j := 0, len(ans)-1; i < j; i, j = i+1, j-1 {
				ans[i], ans[j] = ans[j], ans[i]
			}
		default:
			ans = append(ans, ch)
		}
	}

	return string(ans)
}

func Test_process_string_with_special_operations_i(t *testing.T) {
	tests := []struct {
		s string
	}{
		{"a#b%*"},
		{"z*#"},
	}
	for _, tt := range tests {
		t.Log(processStr(tt.s))
	}
}
