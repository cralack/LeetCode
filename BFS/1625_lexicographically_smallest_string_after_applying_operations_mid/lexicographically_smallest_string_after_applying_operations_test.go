package lexicographicallysmalleststringafterapplyingoperations

import (
	"fmt"
	"testing"
)

func findLexSmallestString(s string, a int, b int) string {
	que := []string{s}
	vis := map[string]bool{s: true}
	ans := s
	n := len(s)
	for len(que) > 0 {
		s = que[0]
		que = que[1:]
		if ans > s {
			ans = s
		}
		t1 := []byte(s)
		for i := 1; i < n; i += 2 {
			t1[i] = byte((int(t1[i]-'0')+a)%10 + '0')
		}
		t2 := s[n-b:] + s[:n-b]
		for _, t := range []string{string(t1), t2} {
			if !vis[t] {
				vis[t] = true
				que = append(que, t)
			}
		}
	}
	return ans
}

func Test_lexicographically_smallest_string(t *testing.T) {
	tests := []struct {
		s    string
		a    int
		b    int
		want string
	}{
		{"5525", 9, 2, "2050"},
		{"74", 5, 1, "24"},
		{"0011", 4, 2, "0011"},
		{"43987654", 7, 3, "00553311"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v %v", tt.s, tt.a, tt.b), func(t *testing.T) {
			if got := findLexSmallestString(tt.s, tt.a, tt.b); got != tt.want {
				t.Errorf("findLexSmallestString() = %v, want %v", got, tt.want)
			}
		})
	}
}
