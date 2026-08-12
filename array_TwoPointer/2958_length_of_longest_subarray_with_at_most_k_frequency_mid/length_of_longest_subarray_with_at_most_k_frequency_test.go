package length_of_longest_subarray_with_at_most_k_frequency_mid

import (
	"testing"
)

func maxSubarrayLength(nums []int, k int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, x := range nums {
		cnt[x]++
		for cnt[x] > k {
			cnt[nums[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}

func Test_length_of_longest_subarray_with_at_most_k_frequency(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
	}{
		{nums: []int{1, 2, 3, 1, 2, 3, 1, 2}, k: 2},
		{nums: []int{1, 2, 1, 2, 1, 2, 1, 2}, k: 1},
		{nums: []int{5, 5, 5, 5, 5, 5, 5}, k: 4},
		{nums: []int{1, 4, 4, 3}, k: 1},
	}
	for _, test := range tests {
		t.Log(maxSubarrayLength(test.nums, test.k))
	}
}
