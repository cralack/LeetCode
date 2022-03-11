package rangesumquery2dimmutable

import (
	"testing"
)

type NumMatrix struct {
	DPtable [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	//构造 一次O(mn)
	row, col := len(matrix), len(matrix[0])
	dp := make([][]int, row+1)
	dp[0] = make([]int, col+1)

	for i, nums := range matrix {
		dp[i+1] = make([]int, col+1)
		for j, v := range nums {
			//上左两排base,dp[i+1][j+1]=上+左-左上+本
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + v
		}
	}
	return NumMatrix{DPtable: dp}
}

func (p *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	//O(1)
	res := p.DPtable[row2+1][col2+1] - p.DPtable[row2+1][col1] -
		p.DPtable[row1][col2+1] + p.DPtable[row1][col1]
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
func Test_range_sum_query_2d_immutable(t *testing.T) {
	// c1 := []string{"NumMatrix", "sumRegion", "sumRegion", "sumRegion"}
	c := [][]int{{2, 1, 4, 3}, {1, 1, 2, 2}, {1, 2, 2, 4}}
	matrix := [][]int{
		{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5}}
	nMatrix := Constructor(matrix)
	for _, arr := range c {
		t.Log(nMatrix.SumRegion(arr[0], arr[1], arr[2], arr[3]))
	}
}
