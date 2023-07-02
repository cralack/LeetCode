package slidingwindowmaximum

import (
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	res := []int{}
	// 单调队列
	window := []int{}
	for i, num := range nums {

		// pop
		if k <= i && 0 < len(window) && nums[i-k] == window[0] {
			window = window[1:]
		}
		// 保持队列内max元素永远在队首&&push
		for 0 < len(window) && window[len(window)-1] < num {
			window = window[:len(window)-1]
		}
		// push
		window = append(window, num)
		// 遍历过k个元素后开始更新res
		if i >= k-1 {
			res = append(res, window[0])
		}
	}
	return res
}
func Test_sliding_window_maximum(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	t.Log(maxSlidingWindow(nums, k))
	nums = []int{1, -1}
	k = 1
	t.Log(maxSlidingWindow(nums, k))
	nums = []int{1, 3, 1, 2, 0, 5}
	k = 3
	t.Log(maxSlidingWindow(nums, k))
}
