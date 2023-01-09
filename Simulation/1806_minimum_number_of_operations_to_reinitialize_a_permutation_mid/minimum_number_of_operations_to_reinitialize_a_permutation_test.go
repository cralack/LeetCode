package minimumnumberofoperations

import (
	"testing"
)

func reinitializePermutation(n int) (step int) {
	perm := make([]int, n)
	for i := range perm {
		perm[i] = i
	}
	tar := append([]int(nil), perm...)
	for {
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				arr[i] = perm[i/2]
			} else {
				arr[i] = perm[n/2+(i-1)/2]
			}
		}
		perm = arr
		step++
		if check(tar, perm) {
			break
		}
	}
	return
}
func check(arr1, arr2 []int) bool {
	n := len(arr1)
	for i := 0; i < n; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func reinitializePermutation_graceful(n int) (ans int) {
	for i := 1; ; {
		ans++
		if i < (n >> 1) {
			i <<= 1
		} else {
			i = (i-(n>>1))<<1 | 1
		}
		if i == 1 {
			return ans
		}
	}
}

func Test_minimum_number_of_operations(t *testing.T) {
	n := 2
	t.Log(reinitializePermutation(n))
	n = 4
	t.Log(reinitializePermutation_graceful(n))
	n = 6
	t.Log(reinitializePermutation(n))
}
