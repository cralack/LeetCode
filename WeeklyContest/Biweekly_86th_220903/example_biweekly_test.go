package biweekly_contest

import (
	"testing"
)

/************ 1st test************/
func findSubarrays(nums []int) bool {
	cntSum := make(map[int]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i] + nums[i+1]
		if _, has := cntSum[sum]; !has {
			cntSum[sum]++
		} else {
			return true
		}
	}
	return false
}
func Test_1st(t *testing.T) {
	nums := []int{4, 2, 4}
	t.Log(findSubarrays(nums))
	nums = []int{1, 2, 3, 4, 5}
	t.Log(findSubarrays(nums))
	nums = []int{0, 0, 0}
	t.Log(findSubarrays(nums))
	nums = []int{1, 2, 3, 1, 0, 2}
	t.Log(findSubarrays(nums))

}

/************ 2nd test************/
func isStrictlyPalindromic(n int) bool {
	for i := 2; i <= n-2; i++ {
		num := turn2knum(n, i)
		if !isPalin(num) {
			return false
		}
	}
	return true
}
func turn2knum(n, k int) (ans []int) {
	for n > 0 {
		ans = append(ans, n%k)
		n /= k
	}
	return
}
func isPalin(num []int) bool {
	for left, right := 0, len(num)-1; left < len(num)/2; left,
		right = left+1, right-1 {
		if num[left] != num[right] {
			return false
		}
	}
	return true
}
func Test_2nd(t *testing.T) {
	n := 9
	t.Log(isStrictlyPalindromic(n))
	n = 4
	t.Log(isStrictlyPalindromic(n))

}

/************ 3rd test************/
func maximumRows(mat [][]int, cols int) (ans int) {
	n, m := len(mat), len(mat[0])
	for i := 0; i < 1<<m; i++ {
		cnt := 0
		for j := 0; j < m; j++ {
			tmp := i >> j & 1
			cnt += tmp
		}
		if cnt != cols {
			continue
		}
		// 统计被覆盖的行
		cnt = 0
		for ii := 0; ii < n; ii++ {
			for jj := 0; jj < m; jj++ {
				if mat[ii][jj] == 1 && i>>jj&1 == 0 {
					goto FAILED
				}
			}
			cnt++
		FAILED:
			continue
		}
		if ans < cnt {
			ans = cnt
		}
	}
	return
}

func Test_3rd(t *testing.T) {
	mat := [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 0, 1}}
	cols := 2
	t.Log(maximumRows(mat, cols))
}

/************ 4th test************/
func maximumRobots(chargeTimes []int, runningCosts []int, budget int) (ans int) {
	type pair struct {
		x, y int
	}
	n := len(chargeTimes)
	que := []pair{} //单调队列维护 max(chargeTimes)
	sum := 0

	for right, left := 0, 0; right < n; right++ {
		if len(que) == 0 {
			left = right
		}

		//末尾元素的chargeTimes小于新进元素
		for len(que) > 0 && que[len(que)-1].x <= chargeTimes[right] {
			que = que[:len(que)-1] //popback
		}
		que = append(que, pair{chargeTimes[right], right})
		sum += runningCosts[right]

		//增大区间直到费用超标
		for len(que) > 0 && que[0].x+(right-left+1)*sum > budget {
			if left == que[0].y {
				que = que[1:] //popfront
			}
			sum -= runningCosts[left]
			left++
		}
		//更新最大区间宽度
		if len(que) > 0 && ans < right-left+1 {
			ans = right - left + 1
		}
	}
	return ans
}

func Test_4th(t *testing.T) {
	chargeTimes := []int{3, 6, 1, 3, 4}
	runningCosts := []int{2, 1, 3, 4, 5}
	budget := 25
	t.Log(maximumRobots(chargeTimes, runningCosts, budget))
	chargeTimes = []int{11, 12, 19}
	runningCosts = []int{210, 8, 7}
	budget = 19
	t.Log(maximumRobots(chargeTimes, runningCosts, budget))
}
