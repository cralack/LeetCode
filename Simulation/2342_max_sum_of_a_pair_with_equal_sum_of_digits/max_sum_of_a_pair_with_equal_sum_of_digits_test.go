package max_sum_of_a_pair_with_equal_sum_of_digits

import (
	"testing"
)

func maximumSum(nums []int) int {
	ans := -1
	dic := make([]int, 82)
	for _, n := range nums {
		key := 0
		for x := n; x > 0; x /= 10 {
			key += x % 10
		}
		if dic[key] > 0 {
			ans = max(ans, dic[key]+n)
		}
		dic[key] = max(dic[key], n)
	}
	return ans
}

func Test_max_sum_of_a_pair_with_equal_sum_of_digits(t *testing.T) {
	tests := []struct{ nums []int }{
		{[]int{18, 43, 36, 13, 7}},
		{[]int{10, 12, 19, 14}},
	}
	for _, tt := range tests {
		t.Log(maximumSum(tt.nums))
	}
}
