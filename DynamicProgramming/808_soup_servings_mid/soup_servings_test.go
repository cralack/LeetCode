package soupservings

import "testing"

func soupServings(n int) float64 {
	if n > 4800 {
		return 1
	}
	f := [200][200]float64{}
	var dfs func(i, j int) float64
	dfs = func(i, j int) float64 {
		// 表示两种汤都分配完了
		if i <= 0 && j <= 0 {
			return 0.5
		} // 当i≤0时，表示汤A先分配完了
		if i <= 0 {
			return 1.0
		} // 当j≤0时，表示汤B先分配完了
		if j <= 0 {
			return 0
		} // 记忆化搜索
		if f[i][j] > 0 {
			return f[i][j]
		}
		// 每一种选择的概率都是0.250.25
		ans := 0.25 * (dfs(i-4, j) + dfs(i-3, j-1) +
			dfs(i-2, j-2) + dfs(i-1, j-3))
		f[i][j] = ans
		return ans
	}
	// 将每25ml的汤视为一份
	return dfs((n+24)/25, (n+24)/25)
}
func Test_soup_servings(t *testing.T) {
	n := 50
	t.Log(soupServings(n))
}
