package partitionarrayintodisjointintervals

import "testing"

func partitionDisjoint_nspace(nums []int) int {
	n := len(nums)
	minArr := make([]int, n)
	minArr[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		minArr[i] = min(minArr[i+1], nums[i])
	}
	max := nums[0]
	for i := 1; ; i++ {
		if max <= minArr[i] {
			return i
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func partitionDisjoint(nums []int) int {
	n := len(nums)
	leftMax, leftPos, curMax := nums[0], 0, nums[0]
	for i := 1; i < n-1; i++ {
		curMax = max(curMax, nums[i])
		if nums[i] < leftMax {
			leftMax = curMax
			leftPos = i
		}
	}
	return leftPos + 1
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func Test_partition_array_into_disjoint_intervals(t *testing.T) {
	nums := []int{5, 0, 3, 8, 6}
	t.Log(partitionDisjoint(nums))
	nums = []int{1, 1, 1, 0, 6, 12}
	t.Log(partitionDisjoint(nums))
}
