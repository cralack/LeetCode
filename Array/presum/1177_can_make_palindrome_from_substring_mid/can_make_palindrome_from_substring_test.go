package canmakepalindromefromsubstring

import (
	"math/bits"
	"testing"
)

func canMakePaliQueries(s string, queries [][]int) []bool {
	//sum[i]保存s[0~i]各个字母出现次数的奇偶性
	sum := make([]uint32, len(s)+1) // 节省一半空间
	for i, c := range s {
		bit := uint32(1) << (c - 'a')
		sum[i+1] = sum[i] ^ bit // 该比特对应字母的奇偶性：奇数变偶数，偶数变奇数
	}

	ans := make([]bool, len(queries))
	for i, q := range queries {
		left, right, k := q[0], q[1], q[2]
		//s[left~right]出现奇数次字母的个数
		m := bits.OnesCount32(sum[right+1] ^ sum[left])
		ans[i] = m/2 <= k
	}
	return ans
}
func Test_can_make_palindrome_from_substring(t *testing.T) {
	tests := []struct {
		s       string
		queries [][]int
	}{
		{s: "abcda", queries: [][]int{{3, 3, 0}, {1, 2, 0},
			{0, 3, 1}, {0, 3, 2}, {0, 4, 1}}},
	}
	for _, tt := range tests {
		t.Log(canMakePaliQueries(tt.s, tt.queries))
	}
}
