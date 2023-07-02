package minimumlimitofballsinabag

import (
	"testing"
)

func minimumSize(nums []int, maxOperations int) int {
	max := nums[0]
	for _, n := range nums {
		if max < n {
			max = n
		}
	}

	left, right := 1, max
	for left < right {
		mid := left + (right-left)>>1
		// 假设每袋放mid个球,总共会消耗sum次操作数
		sum := 0
		for _, n := range nums {
			// 向下取整
			sum += (n - 1) / mid
		}
		if sum <= maxOperations {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func Test_minimum_limit_of_balls_in_a_bag(t *testing.T) {
	nums, max := []int{9}, 2
	t.Log(minimumSize(nums, max))
	nums, max = []int{2, 4, 8, 2}, 4
	t.Log(minimumSize(nums, max))
	nums, max = []int{7, 17}, 2
	t.Log(minimumSize(nums, max))
}
