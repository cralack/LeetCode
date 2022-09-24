package rangesumqueryimmutable

import (
	"testing"
)

type NumArray struct {
	preSum []int
}

func Constructor(nums []int) NumArray {
	pSum := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		pSum[i+1] = pSum[i] + nums[i]
	}
	return NumArray{
		preSum: pSum,
	}
}

func (p *NumArray) SumRange(left int, right int) int {
	return p.preSum[right+1] - p.preSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
func Test_range_sum_query_immutable(t *testing.T) {
	c1 := []string{"NumArray", "sumRange", "sumRange", "sumRange"}
	c2 := [][]int{{-2, 0, 3, -5, 2, -1}, {0, 2}, {2, 5}, {0, 5}}
	var nArr NumArray
	for i, c := range c1 {
		switch c {
		case "NumArray":
			nArr = Constructor(c2[i])
			t.Log(nArr)
		case "sumRange":
			t.Log(nArr.SumRange(c2[i][0], c2[i][1]))
		}
	}
}
