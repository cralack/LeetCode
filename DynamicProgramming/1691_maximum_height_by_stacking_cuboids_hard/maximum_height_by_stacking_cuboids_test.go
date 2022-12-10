package maximumheightbystackingcuboids

import (
	"sort"
	"testing"
)

func maxHeight(cuboids [][]int) (ans int) {
	for _, cube := range cuboids {
		sort.Ints(cube)
	}
	sort.Slice(cuboids, func(i, j int) bool {
		a, b := cuboids[i], cuboids[j]
		return a[0] < b[0] || a[0] == b[0] &&
			(a[1] < b[1] || a[1] == b[1] && a[2] < b[2])
	})
	n := len(cuboids)
	dp := make([]int, n)
	for i := range dp {
		for j := 0; j < i; j++ {
			if cuboids[j][1] <= cuboids[i][1] && cuboids[j][2] <= cuboids[i][2] {
				dp[i] = max(dp[i], dp[j])
			}
		}
		dp[i] += cuboids[i][2]
		ans = max(ans, dp[i])
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_maximum_height_by_stacking_cuboids(t *testing.T) {
	cuboids := [][]int{{50, 45, 20}, {95, 37, 53}, {45, 23, 12}}
	t.Log(maxHeight(cuboids))
	cuboids = [][]int{{38, 25, 45}, {76, 35, 3}}
	t.Log(maxHeight(cuboids))
	cuboids = [][]int{{7, 11, 17}, {7, 17, 11}, {11, 7, 17}, {11, 17, 7}, {17, 7, 11}, {17, 11, 7}}
	t.Log(maxHeight(cuboids))
}
