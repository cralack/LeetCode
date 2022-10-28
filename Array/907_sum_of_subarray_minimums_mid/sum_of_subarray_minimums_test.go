package sumofsubarrayminimums

import "testing"

func sumSubarrayMins(arr []int) (ans int) {
	n := len(arr)
	// 左边界 left[i] 为左侧严格小于 arr[i] 的最近元素位置（不存在时为 -1）
	left := make([]int, n)
	st := []int{-1} // 方便赋值 left
	for i, x := range arr {
		for len(st) > 1 && arr[st[len(st)-1]] >= x {
			st = st[:len(st)-1] // 移除无用数据
		}
		left[i] = st[len(st)-1]
		st = append(st, i)
	}

	// 右边界 right[i] 为右侧小于等于 arr[i] 的最近元素位置（不存在时为 n）
	right := make([]int, n)
	st = []int{n} // 方便赋值 right
	for i := n - 1; i >= 0; i-- {
		for len(st) > 1 && arr[st[len(st)-1]] > arr[i] {
			st = st[:len(st)-1] // 移除无用数据
		}
		right[i] = st[len(st)-1]
		st = append(st, i)
	}

	for i, x := range arr {
		ans += x * (i - left[i]) * (right[i] - i) // 累加贡献
	}
	return ans % (1e9 + 7)
}
func Test_sum_of_subarray_minimums(t *testing.T) {
	arr := []int{3, 1, 2, 4}
	t.Log(sumSubarrayMins(arr))
	arr = []int{11, 81, 94, 43, 3}
	t.Log(sumSubarrayMins(arr))
}
