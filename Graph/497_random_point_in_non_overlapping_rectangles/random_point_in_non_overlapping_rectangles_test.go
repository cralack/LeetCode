package randompointinnonoverlappingrectangles

import (
	"math/rand"
	"sort"
	"testing"
)

type Solution struct {
	rects [][]int
	sum   []int
}

func Constructor(rects [][]int) Solution {
	sum := make([]int, len(rects)+1)
	for i, r := range rects {
		a, b, x, y := r[0], r[1], r[2], r[3]
		sum[i+1] = sum[i] + (x-a+1)*(y-b+1)
	}
	return Solution{rects, sum}
}

func (s *Solution) Pick() []int {
	k := rand.Intn(s.sum[len(s.sum)-1])
	// 在排序的整数切片中搜索 k 并返回 Search 指定的索引。
	rectIndex := sort.SearchInts(s.sum, k+1) - 1
	r := s.rects[rectIndex]
	a, b, y := r[0], r[1], r[3]
	// 随机点数对应当前rect中第n点
	// 当前rect中每行有y-b+1列
	da := (k - s.sum[rectIndex]) / (y - b + 1) // x在当前rect中第da行
	db := (k - s.sum[rectIndex]) % (y - b + 1) // x在当前rect中第db列
	return []int{a + da, b + db}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(rects);
 * param_1 := obj.Pick();
 */
func Test_random_point_in_non_overlapping_rectangles(t *testing.T) {
	c1 := []string{"Solution", "pick", "pick", "pick", "pick", "pick"}
	rects := [][]int{{-2, -2, 1, 1}, {2, 2, 4, 6}}
	var sol Solution
	for _, c := range c1 {
		switch c {
		case "Solution":
			sol = Constructor(rects)
		case "pick":
			t.Log(sol.Pick())
		}
	}
}
