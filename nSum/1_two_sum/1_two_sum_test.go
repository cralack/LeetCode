package twosum

import (
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := [][]int{
		{2, 3, 5, 11, 17, 13, 19, 7, 23, 29},
		{3, 2, 4}, {3, 3}}
	tar := []int{9, 6, 6}
	// for idx, num := range nums {
	// 	t.Log(TwoSum(num, tar[idx]))
	// }

	for idx, num := range nums {
		// t.Log(TwoSumHash(num, tar[idx]))
		ans := TwoSum_v2(num, tar[idx])
		t.Log(ans)

	}
}

func TwoSum(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func TwoSumHash(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func TwoSum_v2(nums []int, target int) (res [][]int) {
	sort.Ints(nums)
	lo, hi := 0, len(nums)-1
	for lo < hi {
		sum := nums[lo] + nums[hi]
		left, right := nums[lo], nums[hi]
		if sum < target {
			for lo < hi && nums[lo] == left {
				lo++
			}
		} else if sum > target {
			for lo < hi && nums[hi] == right {
				hi--
			}
		} else if sum == target {
			res = append(res, []int{left, right})
			for lo < hi && nums[lo] == left {
				lo++
			}
			for lo < hi && nums[hi] == right {
				hi--
			}
		}
	}
	return res
}
