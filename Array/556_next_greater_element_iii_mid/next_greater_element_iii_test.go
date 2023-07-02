package nextgreaterelementiiimid

import (
	"math"
	"testing"
)

func nextGreaterElement(n int) int {
	arr := []int{}
	for n != 0 {
		arr = append(arr, n%10)
		n /= 10
	}
	n = len(arr)
	idx := -1
	// find out first ↗
	for i := 0; i < n-1 && idx == -1; i++ {
		if arr[i+1] < arr[i] {
			idx = i + 1
		}
	}
	if idx == -1 {
		return idx
	}
	// swap target
	for i := 0; i < idx; i++ {
		if arr[i] > arr[idx] {
			arr[i], arr[idx] = arr[idx], arr[i]
			break
		}
	}
	// reverse
	l, r := 0, idx-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
	// arr to int
	ans := 0
	for i := n - 1; i >= 0; i-- {
		ans = ans*10 + arr[i]
		if ans > math.MaxInt32 {
			return -1
		}
	}
	return ans
}

func Test_next_greater_element_iii(t *testing.T) {
	n := 12
	t.Log(nextGreaterElement(n))
	n = 21
	t.Log(nextGreaterElement(n))
	n = 20220703
	t.Log(nextGreaterElement(n))
	n = 12321
	t.Log(nextGreaterElement(n))
}
