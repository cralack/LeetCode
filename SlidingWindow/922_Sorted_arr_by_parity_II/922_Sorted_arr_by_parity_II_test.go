package sortedarrbyparityii

import (
	"testing"
)

func sortArrayByParityII(nums []int) []int {
	for even, odd := 0, 1; even < len(nums); even += 2 {
		if nums[even]%2 == 1 {
			for nums[odd]%2 == 1 {
				odd += 2
			}
			nums[even], nums[odd] = nums[odd], nums[even]
		}
	}

	return nums
}
func TestSortArrByParityII(t *testing.T) {
	arrs := [][]int{{4, 2, 5, 7}, {2, 3}}
	ans := [][]int{}
	for _, arr := range arrs {
		ans = append(ans, sortArrayByParityII(arr))
	}
	t.Log(ans)
}
