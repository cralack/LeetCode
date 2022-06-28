package wigglesortiimid

import (
	"sort"
	"testing"
)

func wiggleSort(nums []int) {
	n := len(nums)
	tmp := make([]int, n)
	j, k := n, (n+1)>>1
	copy(tmp, nums)
	sort.Ints(tmp)
	for i := 0; i < n; i++ {
		if i&1 == 1 {
			j--
			nums[i] = tmp[j]
		} else {
			k--
			nums[i] = tmp[k]
		}
	}
}
func Test_wiggle_sort_ii(t *testing.T) {
	nums := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums)
	t.Log(nums)
	nums = []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums)
	t.Log(nums)
	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	wiggleSort(nums)
	t.Log(nums)

}
