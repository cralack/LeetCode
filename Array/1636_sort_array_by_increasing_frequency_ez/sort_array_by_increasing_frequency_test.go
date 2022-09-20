package sortarraybyincreasingfrequencyez

import (
	"sort"
	"testing"
)

func frequencySort(nums []int) []int {
	freq := make(map[int]int, 0)
	for _, n := range nums {
		freq[n]++
	}

	sort.Slice(nums, func(i, j int) bool {
		if freq[nums[i]] == freq[nums[j]] {
			return nums[i] > nums[j]
		}
		return freq[nums[i]] < freq[nums[j]]
	})
	return nums
}
func Test_sort_array_by_increasing_frequency(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3}
	t.Log(frequencySort(nums))
	nums = []int{2, 3, 1, 3, 2}
	t.Log(frequencySort(nums))
	nums = []int{-1, 1, -6, 4, 5, -6, 1, 4, 1}
	t.Log(frequencySort(nums))
}
