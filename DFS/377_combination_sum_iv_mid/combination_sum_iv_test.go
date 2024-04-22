package combination_sum_iv_mid

import (
	"testing"
)

func combinationSum4DP(nums []int, target int) int {
	f := make([]int, target+1)
	f[0] = 1
	for i := 1; i <= target; i++ {
		for _, x := range nums {
			if x <= i {
				f[i] += f[i-x]
			}
		}
	}
	return f[target]
}

func combinationSum4(nums []int, target int) int {
	memo := make([]int, target+1)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int) int
	dfs = func(index int) (res int) {
		if index == 0 {
			return 1
		}
		p := &memo[index]
		if *p != -1 {
			return *p
		}
		for _, x := range nums {
			if x <= index {
				res += dfs(index - x)
			}
		}
		*p = res
		return
	}
	return dfs(target)
}

func Test_combination_sum_iv(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
	}{
		{nums: []int{1, 2, 3}, target: 4},
		{nums: []int{9}, target: 3},
		{nums: []int{4, 2, 1}, target: 32},
	}
	for _, test := range tests {
		t.Log(combinationSum4(test.nums, test.target))
		t.Log(combinationSum4DP(test.nums, test.target))
	}
}
