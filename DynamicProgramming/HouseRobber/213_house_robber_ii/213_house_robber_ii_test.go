package houserobberii

import "testing"

func rob(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	rob_range := func(start, end int) int {
		dp2, dp1 := 0, 0
		dpi := 0
		for i := start; i < end; i++ {
			dpi = max(dp1, dp2+nums[i])
			dp2 = dp1
			dp1 = dpi
		}
		return dpi
	}
	return max(rob_range(0, size-1), rob_range(1, size))
}
func Test_house_robber_ii(t *testing.T) {
	streets := [][]int{{2, 3, 2}, {1, 2, 3, 1}, {1, 2, 3}, {1}}
	for _, houses := range streets {
		res := rob(houses)
		t.Log(res)
	}
}
