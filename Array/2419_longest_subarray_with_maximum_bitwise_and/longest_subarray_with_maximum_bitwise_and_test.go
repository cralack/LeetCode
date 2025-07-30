package longestsubarraywithmaximumbitwiseand

import "testing"

func longestSubarray(nums []int) (ans int) {
	mx, cnt := nums[0], 0
	for _, x := range nums {
		if x > mx {
			mx = x
			cnt = 1
			ans = 1
		} else if x == mx {
			cnt++
			ans = max(ans, cnt)
		} else {
			cnt = 0
		}
	}
	return
}

func Test_longest_subarray_with_maximum_bitwise_and(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{1, 2, 3, 3, 2, 2}, want: 2},
		{nums: []int{1, 2, 3, 4}, want: 1},
		{nums: []int{96317, 96317, 96317, 96317, 96317,
			96317, 96317, 96317, 96317, 279979}, want: 1},
		{nums: []int{311155, 311155, 311155, 311155, 311155,
			311155, 311155, 311155, 201191, 311155}, want: 8},
	}
	for _, tt := range tests {
		t.Log(longestSubarray(tt.nums) == tt.want)
	}
}
