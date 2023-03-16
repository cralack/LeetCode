package countsubarrayswithmediank

import "testing"

func countSubarrays(nums []int, k int) int {
	n, x := len(nums), len(nums)
	mp := make(map[int]int, 0)
	mp[0] = 1
	preSum, res := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] == k { //与k相等的数字为x
			preSum += x
		} else if nums[i] < k { //记比k小的数字为-1
			preSum--
		} else { //记比k大的数字为1
			preSum++
		}
		//数组长度为奇数：子数组和为x
		if val, has := mp[preSum-x]; has {
			res += val
		}
		////数组长度为偶数：子数组和为x+1
		if val, has := mp[preSum-x-1]; has {
			res += val
		}
		mp[preSum]++
	}
	return res
}

func Test_count_subarrays_with_median_k(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
	}{
		{[]int{3, 2, 1, 4, 5}, 4},
		{[]int{2, 3, 1}, 3},
	}
	for _, tt := range tests {
		t.Log(countSubarrays(tt.nums, tt.k))
	}
}
