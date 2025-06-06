package using_a_robot_to_print_the_lexicographically_smallest_string_mid

import (
	"fmt"
	"testing"
)

func robotWithString(s string) string {
	n := len(s)
	sufMin := make([]byte, n+1) // 记录后缀最小值
	sufMin[n] = byte(0xff)      // 哨兵
	for i := n - 1; i >= 0; i-- {
		sufMin[i] = min(sufMin[i+1], s[i])
	}
	ans := make([]byte, 0, n)
	stack := make([]byte, 0, n)
	for i, char := range s {
		stack = append(stack, byte(char))
		// 栈顶与后缀相比，栈顶小则出栈入ans
		for len(stack) > 0 && stack[len(stack)-1] <= sufMin[i+1] {
			ans = append(ans, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}
	return string(ans)
}

func Test_using_a_robot_to_print_the_lexicographically_smallest_string(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{s: "zza", want: "azz"},
		{s: "bac", want: "abc"},
		{s: "bdda", want: "addb"},
		{s: "caba", want: "aabc"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got := robotWithString(tt.s)
			if got != tt.want {
				t.Errorf("\ngot: %v\nwant: %v", got, tt.want)
			}
		})
	}
}
