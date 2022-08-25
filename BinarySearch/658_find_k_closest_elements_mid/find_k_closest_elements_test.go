package findkclosestelementsmid

import (
	"testing"
)

func findClosestElements(arr []int, k int, x int) []int {
	left := 0
	right := len(arr) - k - 1
	for left <= right {
		mid := left + (right-left)/2
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return arr[left : left+k]
}

func Test_find_k_closest_elements(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k := 4
	x := 3
	t.Log(findClosestElements(arr, k, x))
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k = 4
	x = 11
	t.Log(findClosestElements(arr, k, x))
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	k = 4
	x = -1
	t.Log(findClosestElements(arr, k, x))
}
