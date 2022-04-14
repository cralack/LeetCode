package intervallistintersections

import "testing"

func intervalIntersection(firstList [][]int,
	secondList [][]int) (res [][]int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	n, m := len(firstList), len(secondList)
	i, j := 0, 0

	for i < n && j < m {
		l1, r1 := firstList[i][0], firstList[i][1]
		l2, r2 := secondList[j][0], secondList[j][1]
		if r2 >= l1 && r1 >= l2 { //两个区间存在交集
			res = append(res, []int{max(l1, l2), min(r1, r2)})
		}
		if r2 < r1 {
			j++
		} else {
			i++
		}
	}
	return
}
func Test_interval_list_intersections(t *testing.T) {
	firstList := [][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}}
	secondList := [][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}}
	t.Log(intervalIntersection(firstList, secondList))
}
