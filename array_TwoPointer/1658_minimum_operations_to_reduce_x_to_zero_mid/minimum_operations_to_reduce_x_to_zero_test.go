package minimumoperations

import "testing"

func minOperations(nums []int, x int) (ans int) {
	n, sum := len(nums), 0
	right := n
	// 计算最长后缀
	for right > 0 && sum+nums[right-1] <= x {
		right--
		sum += nums[right]
	}
	if right == 0 && sum < x {
		return -1
	}
	ans = n + 1
	if sum == x {
		ans = n - right
	}
	for left, num := range nums {
		sum += num
		// 缩小后缀长度
		for ; right < n && sum > x; right++ {
			sum -= nums[right]
		} // 缩小失败，说明前缀过长
		if sum > x {
			break
		}
		if sum == x { // 前缀+后缀长度
			ans = min(ans, left+1+n-right)
		}
	}
	if ans > n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_minimum_operations_to_reduce_x_to_zero(t *testing.T) {
	nums, x := []int{1, 1, 4, 2, 3}, 5
	t.Log(minOperations(nums, x))
	nums, x = []int{5, 6, 7, 8, 9}, 4
	t.Log(minOperations(nums, x))
	nums, x = []int{3, 2, 20, 1, 1, 3}, 10
	t.Log(minOperations(nums, x))
}
