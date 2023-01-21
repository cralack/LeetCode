package minimumsidewayjumps

import "testing"

func minSideJumps(obstacles []int) int {
	n := len(obstacles)
	dp := make([][3]int, n)
	dp[0][0] = 1
	dp[0][1] = 0
	dp[0][2] = 1
	for i := 1; i < n; i++ {
		v := obstacles[i] - 1
		minCost := 1 << 30
		for j := 0; j < 3; j++ {
			if v == j {
				dp[i][j] = 1 << 30
			} else {
				minCost = min(minCost, dp[i-1][j])
			}
		}
		for j := 0; j < 3; j++ {
			if v != j {
				dp[i][j] = min(dp[i-1][j], minCost+1)
			}
		}
	}
	n--
	return min(dp[n][0], min(dp[n][1], dp[n][2]))
}

func minSideJumps_1dshrink(obstacles []int) int {
	dp := [3]int{1, 0, 1}
	const inf = 1 << 30
	for _, v := range obstacles[1:] {
		for j := 0; j < 3; j++ {
			if v == j+1 {
				dp[j] = inf
				break
			}
		}
		cost := min(dp[0], min(dp[1], dp[2])) + 1
		for j := 0; j < 3; j++ {
			if v != j+1 {
				dp[j] = min(dp[j], cost)
			}
		}
	}
	return min(dp[0], min(dp[1], dp[2]))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func Test_minimum_sideway_jumps(t *testing.T) {
	obstacles := []int{0, 1, 2, 3, 0}
	t.Log(minSideJumps(obstacles))
	t.Log(minSideJumps_1dshrink(obstacles))
	obstacles = []int{0, 1, 1, 3, 3, 0}
	t.Log(minSideJumps(obstacles))
	t.Log(minSideJumps_1dshrink(obstacles))
	obstacles = []int{0, 2, 1, 0, 3, 0}
	t.Log(minSideJumps(obstacles))
	t.Log(minSideJumps_1dshrink(obstacles))
}
