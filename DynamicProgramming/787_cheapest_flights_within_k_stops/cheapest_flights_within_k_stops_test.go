package cheapestflightswithinkstops

import (
	"math"
	"testing"
)

func min(a ...int) int {
	res := a[0]
	for _, v := range a[1:] {
		if res > v {
			res = v
		}
	}
	return res
}
func findCheapestPrice_rec(n int, flights [][]int, src int, dst int, k int) int {
	k++
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, k+1)
		for j := range memo[i] {
			memo[i][j] = -888
		}
	}
	// 将中转站个数转化成边的条数
	indegree := make(map[int][][]int, 0)
	for _, f := range flights {
		from, to, price := f[0], f[1], f[2]
		indegree[to] = append(indegree[to], []int{from, price})
	}
	var dp func(s, k int) int
	// 定义：从 src 出发，k 步之内到达 s 的最短路径权重
	dp = func(s, k int) int {
		// base case
		if s == src {
			return 0
		}
		if k == 0 {
			return -1
		}
		if memo[s][k] != -888 {
			return memo[s][k]
		}
		// 初始化为最大值，方便等会取最小值
		res := math.MaxInt32
		if _, ok := indegree[s]; ok {
			for _, v := range indegree[s] {
				from := v[0]
				price := v[1]
				subProblem := dp(from, k-1)
				if subProblem != -1 {
					res = min(res, subProblem+price)
				}
			}
		}
		if res == math.MaxInt32 {
			memo[s][k] = -1
		} else {
			memo[s][k] = res
		}
		return memo[s][k]
	}
	return dp(dst, k)
}
func findCheapestPrice_ite(n int, flights [][]int, src int, dst int, k int) int {
	const INF = math.MaxInt32
	dp := make([][]int, k+2)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}
	dp[0][src] = 0
	for times := 1; times <= k+1; times++ {
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			dp[times][to] = min(dp[times][to], dp[times-1][from]+price)
		}
	}
	ans := INF
	for t := 1; t <= k+1; t++ {
		ans = min(ans, dp[t][dst])
	}
	if ans == INF {
		ans = -1
	}
	return ans
}
func Test_cheapest_flights_within_k_stops(t *testing.T) {
	n, edges := 3, [][]int{{0, 1, 100}, {1, 2, 100}, {0, 2, 500}}
	src, dst, k := 0, 2, 0
	t.Log(findCheapestPrice_rec(n, edges, src, dst, k))
	t.Log(findCheapestPrice_ite(n, edges, src, dst, k))
}
