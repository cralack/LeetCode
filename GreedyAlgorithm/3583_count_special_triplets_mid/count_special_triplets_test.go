package count_special_triplets_mid

import (
	"fmt"
	"sort"
	"testing"
)

func specialTriplets(nums []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(nums)
	dic := make(map[int][]int)
	for idx, num := range nums {
		dic[num] = append(dic[num], idx)
	}
	for i := 1; i < n-1; i++ {
		num := nums[i]
		if ids, has := dic[num*2]; has {
			left := sort.SearchInts(ids, i)
			right := len(ids) - left
			if left < len(ids) && ids[left] == i {
				right--
			}
			ans += left * right % mod
		}
	}
	return ans % mod
}

// damn bro,its really elegant
func specialTripletsV2(nums []int) (ans int) {
	const mod int = 1e9 + 7
	suf := map[int]int{}
	for _, n := range nums {
		suf[n]++
	}
	pre := map[int]int{}
	for _, x := range nums {
		suf[x]--
		ans += pre[x*2] * suf[x*2]
		pre[x]++
	}
	return ans % mod
}

func Test_count_special_triplets(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{nums: []int{6, 3, 6}, want: 1},
		{nums: []int{0, 1, 0, 0}, want: 1},
		{nums: []int{8, 4, 2, 8, 4}, want: 2},
	}
	for i, test := range tests {
		t.Run(fmt.Sprintf("count:%d", i), func(t *testing.T) {
			t.Log(specialTriplets(test.nums) == test.want)
			t.Log(specialTripletsV2(test.nums) == test.want)
		})
	}
}
