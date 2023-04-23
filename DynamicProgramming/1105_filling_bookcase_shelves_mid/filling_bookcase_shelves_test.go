package fillingbookcaseshelves

import (
	"math"
	"testing"
)

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	f := make([]int, n+1) // f[0]=0，翻译自 dfs(-1)=0
	for i := range books {
		f[i+1] = math.MaxInt
		maxH, leftW := 0, shelfWidth
		for j := i; j >= 0; j-- {
			leftW -= books[j][0]
			if leftW < 0 { // 空间不足，无法放书
				break
			}
			maxH = max(maxH, books[j][1]) // 从 j 到 i 的最大高度
			f[i+1] = min(f[i+1], f[j]+maxH)
		}
	}
	return f[n] // 翻译自 dfs(n-1)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func Test_filling_bookcase_shelves(t *testing.T) {
	tests := []struct {
		books      [][]int
		shelfWidth int
		wanna      int
	}{
		{[][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4, 6},
		{[][]int{{1, 3}, {2, 4}, {3, 2}}, 6, 4},
	}
	for _, tt := range tests {
		t.Log(minHeightShelves(tt.books, tt.shelfWidth) == tt.wanna)
	}
}
