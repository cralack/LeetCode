package maximumequalfrequencyhard

import "testing"

func maxEqualFreq(nums []int) (ans int) {
	max := 0
	cnt := make(map[int]int, 0)
	freq := make(map[int]int, 0)
	for i, num := range nums {
		if cnt[num] > 0 {
			freq[cnt[num]]--
		}
		cnt[num]++
		freq[cnt[num]]++
		max = maxComp(max, cnt[num])
		// 所有数的出现次数都是一次
		if max == 1 ||
			// 所有数的出现次数都是 maxFreq或maxFreq−1并且最大出现次数的数只有一个
			freq[max]*max+freq[max-1]*(max-1) == i+1 && freq[max] == 1 ||
			// 除开一个数，其他所有数的出现次数都是maxFreq，并且该数的出现次数为 1
			freq[max]*max+1 == i+1 && freq[1] == 1 {
			ans = maxComp(ans, i+1)
		}
	}
	return
}

func maxComp(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_maximum_equal_frequency(t *testing.T) {
	arr := []int{2, 2, 1, 1, 5, 3, 3, 5}
	t.Log(maxEqualFreq(arr))
	arr = []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}
	t.Log(maxEqualFreq(arr))
}
