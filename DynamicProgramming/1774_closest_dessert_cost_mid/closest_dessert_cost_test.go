package closestdessertcost

import "testing"

func closestCost_backtrack(baseCosts []int, toppingCosts []int, target int) (ans int) {
	ans = baseCosts[0]
	var dfs func(int, int)
	dfs = func(idx, curCost int) {
		if abs(ans-target) < curCost-target {
			return
		} else if abs(ans-target) >= abs(curCost-target) {
			if abs(ans-target) > abs(curCost-target) {
				ans = curCost
			} else {
				ans = min(ans, curCost)
			}
		}

		if idx == len(toppingCosts) {
			return
		}
		dfs(idx+1, curCost+toppingCosts[idx]*2)
		dfs(idx+1, curCost+toppingCosts[idx])
		dfs(idx+1, curCost)
	}
	for _, baseCost := range baseCosts {
		dfs(0, baseCost)
	}
	return ans
}

func closestCost_DP(baseCosts []int, toppingCosts []int, target int) int {
	x := baseCosts[0]
	for _, c := range baseCosts {
		x = min(x, c)
	}
	if x > target {
		return x
	}
	can := make([]bool, target+1)
	ans := 2*target - x
	for _, c := range baseCosts {
		if c <= target {
			can[c] = true
		} else {
			ans = min(ans, c)
		}
	}
	for _, c := range toppingCosts {
		for count := 0; count < 2; count++ {
			for i := target; i > 0; i-- {
				if can[i] && i+c > target {
					ans = min(ans, i+c)
				}
				if i-c > 0 {
					can[i] = can[i] || can[i-c]
				}
			}
		}
	}
	for i := 0; i <= ans-target; i++ {
		if can[target-i] {
			return target - i
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_closest_dessert_cost(t *testing.T) {
	baseCost := []int{1, 7}
	toppingCost := []int{3, 4}
	target := 10
	t.Log(closestCost_backtrack(baseCost, toppingCost, target))
	t.Log(closestCost_DP(baseCost, toppingCost, target))

	baseCost = []int{2, 3}
	toppingCost = []int{4, 5, 100}
	target = 18
	t.Log(closestCost_backtrack(baseCost, toppingCost, target))
	t.Log(closestCost_DP(baseCost, toppingCost, target))

	baseCost = []int{3, 10}
	toppingCost = []int{2, 5}
	target = 9
	t.Log(closestCost_backtrack(baseCost, toppingCost, target))
	t.Log(closestCost_DP(baseCost, toppingCost, target))

	baseCost = []int{10}
	toppingCost = []int{1}
	target = 1
	t.Log(closestCost_backtrack(baseCost, toppingCost, target))
	t.Log(closestCost_DP(baseCost, toppingCost, target))
}

func Benchmark_asd(b *testing.B) {
	baseCost := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	toppingCost := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	target := 21
	b.Run("backtrack", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			closestCost_backtrack(baseCost, toppingCost, target)
		}
		b.StopTimer()
	})
	b.Run("DP", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			closestCost_DP(baseCost, toppingCost, target)
		}
		b.StopTimer()
	})
}
