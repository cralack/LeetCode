package dungeongame

import (
	"testing"
)

func calculateMinimumHP_recursion(grid [][]int) int {
	const INT_MAX = int(^uint(0) >> 1)
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res > v {
				res = v
			}
		}
		return res
	}
	// init
	m, n := len(grid), len(grid[0])
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	/* 定义：从 (i, j) 到达右下角，需要的初始血量至少是多少 */
	var dp func(i, j int) int
	dp = func(i, j int) int {
		// base case
		if i == m-1 && j == n-1 { // end
			if grid[i][j] >= 0 {
				return 1
			}
			return -grid[i][j] + 1
		}
		// 触及边界
		if i == m || j == n {
			return INT_MAX
		}
		// 避免重复计算
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		// 状态转移逻辑
		res := min(dp(i, j+1), dp(i+1, j)) - grid[i][j]
		if res <= 0 {
			memo[i][j] = 1
		} else {
			memo[i][j] = res
		}
		return memo[i][j]
	}

	res := dp(0, 0)
	return res
}
func calculateMinimumHP_iterate(grid [][]int) int {
	// init
	const INT_MAX = int(^uint(0) >> 1)
	min := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if res > v {
				res = v
			}
		}
		return res
	}
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	// base case
	if grid[m-1][n-1] < 0 {
		dp[m-1][n-1] = -grid[m-1][n-1] + 1
	} else {
		dp[m-1][n-1] = 1
	}

	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			// 边界
			if i == m || j == n {
				dp[i][j] = INT_MAX
				continue
			}
			// 略过base case
			if i == m-1 && j == n-1 {
				continue
			} // 状态转移逻辑
			res := min(dp[i+1][j], dp[i][j+1]) - grid[i][j]
			if res <= 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = res
			}
		}
	}
	return dp[0][0]
}
func Test_grid_game(t *testing.T) {
	grid := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5},
		{0, -10, 20}, {0, 0, 100}, {-1, -1, 0}}
	t.Log(calculateMinimumHP_iterate(grid))
	grid = [][]int{{2}, {1}}
	t.Log(calculateMinimumHP_recursion(grid))
}

func Benchmark_dungeon_game(b *testing.B) {
	grid := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	b.Run("ite", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			calculateMinimumHP_iterate(grid)
		}
		b.StopTimer()
	})
	b.Run("ruc", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			calculateMinimumHP_recursion(grid)
		}
		b.StopTimer()
	})
}
