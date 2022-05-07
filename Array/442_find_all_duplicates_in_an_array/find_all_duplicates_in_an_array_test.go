package findallduplicatesinanarray

import "testing"

func findDuplicates(nums []int) (res []int) {
	for _, x := range nums {
		if x < 0 {
			x = -x
		}
		if nums[x-1] > 0 {
			nums[x-1] = -nums[x-1]
		} else {
			res = append(res, x)
		}
	}
	return
}
func Test_find_all_duplicates_in_an_array(t *testing.T) {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	t.Log(findDuplicates(nums))
}
