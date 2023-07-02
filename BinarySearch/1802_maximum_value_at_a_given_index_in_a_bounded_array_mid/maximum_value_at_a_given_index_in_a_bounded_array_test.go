package maximumvalue

import (
	"sort"
	"testing"
)

func maxValue(n int, index int, maxSum int) int {
	// 假设山顶nums[index]==x
	// 在index左侧的数组元素从x-1每次递减1,最小到1
	// 在index及右侧的数组元素从x也是每次递减1,最小到1
	sum := func(x, cnt int) int {
		if x >= cnt {
			// 左侧元素之和
			return (x + x - cnt + 1) * cnt / 2
		} // 右侧元素之和
		return (x+1)*x/2 + cnt - x
	}
	return sort.Search(maxSum, func(mid int) bool {
		mid++
		return sum(mid-1, index)+sum(mid, n-index) > maxSum
	})
}

func Test_maximum_value_at_a_given_index_in_a_bounded_array(t *testing.T) {
	n := 4
	index := 2
	maxSum := 6
	t.Log(maxValue(n, index, maxSum))
	n = 6
	index = 1
	maxSum = 10
	t.Log(maxValue(n, index, maxSum))
	n = 1
	index = 0
	maxSum = 24
	t.Log(maxValue(n, index, maxSum))
}
