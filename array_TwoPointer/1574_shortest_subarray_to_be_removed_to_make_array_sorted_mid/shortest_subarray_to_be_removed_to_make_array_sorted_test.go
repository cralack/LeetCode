package shortestsubarraytoberemovedtomakearraysorted

import "testing"

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	right := n - 1
	for right > 0 && arr[right-1] <= arr[right] {
		right--
	}
	if right == 0 { // arr 已经是非递减数组
		return 0
	}
	// 此时 arr[right-1] > arr[right]
	ans := right // 删除 arr[:right]
	for left := 0; left == 0 || arr[left-1] <= arr[left]; left++ {
		for right < n && arr[right] < arr[left] {
			right++
		}
		ans = min(ans, right-left-1) // 删除 arr[left+1:right]
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_shortest_subarray_to_be_removed(t *testing.T) {
	tests := []struct {
		arr   []int
		wanna int
	}{
		{[]int{1, 2, 3, 10, 4, 2, 3, 5}, 3},
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Log(findLengthOfShortestSubarray(tt.arr) == tt.wanna)
	}
}
