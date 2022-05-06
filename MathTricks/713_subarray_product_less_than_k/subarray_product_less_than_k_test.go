package subarrayproductlessthank

import "testing"

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}
func Test_subarray_product_less_than_k(t *testing.T) {
	nums, k := []int{10, 5, 2, 6}, 100
	t.Log(numSubarrayProductLessThanK(nums, k))
	nums, k = []int{1, 2, 3}, 0
	t.Log(numSubarrayProductLessThanK(nums, k))
}
