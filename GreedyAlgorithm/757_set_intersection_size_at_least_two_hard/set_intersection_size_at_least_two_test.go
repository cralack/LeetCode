package setintersectionsizeatleasttwohard

import (
	"sort"
	"testing"
)

func intersectionSizeTwo(intervals [][]int) (ans int) {
	sort.Slice(intervals, func(i, j int) bool {
		a, b := intervals[i], intervals[j]
		// 右端点从小到大，保证贪心处理边缘点的正确性，
		// 同时右端点一样的时候优先处理长度最短的区间，因为该区间可选点最少
		return a[1] < b[1] || a[1] == b[1] && b[0] < a[0]
	})
	// 由于是单调递增添加元素, 维护当前集合前二大的元素即可判断是否需要添加新的
	a, b := -1, -1
	for _, interval := range intervals {
		// 如果区间左端点也在当前最大元素的右边
		// 说明需要从该区间添加两个新点(可以理解为递归，前面的不再起作用)
		if interval[0] > b {
			// 贪心取最大的两个点
			a = interval[1] - 1
			b = interval[1]
			ans += 2
			// 如果区间左端点位于当前最大元素与次大元素的中间
			// 说明最大元素本身是区间内的一个点了，
			// 我们还需要再取一个，保留b付给a同时取该区间最大
		} else if interval[0] > a {
			a = b
			b = interval[1]
			ans++
		}
	}
	return
}
func Test_set_intersection_size_at_least_two(t *testing.T) {
	intervals := [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}
	t.Log(intersectionSizeTwo(intervals))
	intervals = [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	t.Log(intersectionSizeTwo(intervals))
}
