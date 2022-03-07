package threesum

import (
	"sort"
	"testing"
)

func Test_3sum(t *testing.T) {
	arrs := [][]int{{-1, 0, 1, 2, -1, -4}, {}, {0}}
	for _, arr := range arrs {
		res := threeSum(arr)
		t.Log(res)
	}
}
func TwoSum(nums []int, start, target int) (res [][]int) {
	sort.Ints(nums)
	lo, hi := start, len(nums)-1
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
		} else {
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

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	size := len(nums)
	target := 0
	res := [][]int{}
	for i := 0; i < size; i++ {
		tuples := TwoSum(nums, i+1, target-nums[i])
		for _, tuple := range tuples {
			tuple = append(tuple, nums[i])
			res = append(res, tuple)
		}
		for i < size-1 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
