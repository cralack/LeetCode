package maximumfruitsharvestedafteratmostksteps

import (
	"sort"
	"testing"
)

func maxTotalFruits(fruits [][]int, startPos int, k int) (ans int) {
	n := len(fruits)
	// 向左最远能到 fruits[left][0]
	left := sort.Search(n, func(i int) bool {
		return fruits[i][0] >= startPos-k
	})
	right, sum := left, 0

	for ; right < n && fruits[right][0] <= startPos; right++ {
		sum += fruits[right][1] // 从 fruits[left][0] 到 startPos 的水果数
	}
	ans = sum
	for ; right < n && fruits[right][0] <= startPos+k; right++ {
		sum += fruits[right][1]
		//先右后左距离: (fruits[right][0]−startPos)+(fruits[right][0]−fruits[left][0])
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			//先左后右距离:(startPos−fruits[left][0])+(fruits[right][0]−fruits[left][0])
			fruits[right][0]-fruits[left][0]*2+startPos > k {
			// fruits[left][0] 无法到达
			sum -= fruits[left][1]
			//如果上面两个式子均大于 k，就说明 fruits[left][0] 太远了，需要增加left
			left++
		}
		ans = max(ans, sum)
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func Test_maximum_fruits_harvested_after_at_most_k_steps(t *testing.T) {
	tests := []struct {
		fruits   [][]int
		startPos int
		k        int
		wanna    int
	}{
		{[][]int{{2, 8}, {6, 3}, {8, 6}}, 5, 4, 9},
		{[][]int{{0, 9}, {4, 1}, {5, 7}, {6, 2}, {7, 4}, {10, 9}}, 5, 4, 14},
		{[][]int{{0, 3}, {6, 4}, {8, 5}}, 3, 2, 0},
	}

	for _, tt := range tests {
		t.Log(maxTotalFruits(tt.fruits, tt.startPos, tt.k) == tt.wanna)
	}
}
