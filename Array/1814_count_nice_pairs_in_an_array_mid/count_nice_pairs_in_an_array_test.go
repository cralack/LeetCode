package countnicepairsinanarray

import "testing"

func countNicePairs(nums []int) (ans int) {
	const mod int = 1e9 + 7
	//if nums[i]+rev(nums[j])=nums[j]+rev(nums[i])
	//so nums[i]-rev(nums[i])==nums[j]-rev(nums[j])
	cnt := make(map[int]int, len(nums))
	for _, n := range nums {
		cur := n - rev(n)
		ans += cnt[cur]
		ans %= mod
		cnt[cur]++
	}
	return
}

func rev(num int) (ans int) {
	for num > 0 {
		ans *= 10
		ans += num % 10
		num /= 10
	}
	return ans
}
func Test_count_nice_pairs_in_an_array(t *testing.T) {
	nums := []int{42, 11, 1, 97}
	t.Log(countNicePairs(nums))
	nums = []int{13, 10, 35, 24, 76}
	t.Log(countNicePairs(nums))
}
