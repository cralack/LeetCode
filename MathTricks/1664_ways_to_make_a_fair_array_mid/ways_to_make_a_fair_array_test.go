package waystomakeafairarray

import "testing"

func waysToMakeFair(nums []int) (ans int) {
	// oddSum evenSum 记录平时偶数位奇数位的总和
	evenSum, oddSum := 0, 0
	for i, n := range nums {
		if i%2 == 0 {
			evenSum += n
		} else {
			oddSum += n
		}
	}
	// 记录当前位之前的奇偶数位的总和
	evenCur, oddCur := 0, 0
	for i, cur := range nums {
		// cur下标为偶
		if i%2 == 0 {
			// 奇数下标和(cur前奇数和+cur后偶数和)==偶数下标和(cur前偶数和+cur后奇数和)
			if oddCur+evenSum-evenCur-cur == oddSum+evenCur-oddCur {
				ans++
			} // 更新cur前偶数和
			evenCur += cur
		} else {
			if evenSum+oddCur-evenCur == evenCur+oddSum-oddCur-cur {
				ans++
			}
			oddCur += cur
		}
	}
	return
}
func Test_ways_to_make_a_fair_array(t *testing.T) {
	nums := []int{2, 1, 6, 4}
	t.Log(waysToMakeFair(nums))
	nums = []int{1, 1, 1}
	t.Log(waysToMakeFair(nums))
	nums = []int{1, 2, 3}
	t.Log(waysToMakeFair(nums))
}
