package biweekly_contest

import "testing"

/************ 1st test************/
func separateDigits(nums []int) (ans []int) {
	for _, n := range nums {
		stack := []int{}
		for n > 0 {
			stack = append(stack, n%10)
			n /= 10
		}
		for len(stack) > 0 {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, cur)
		}
	}
	return ans
}
func Test_1st(t *testing.T) {
	nums := []int{13, 25, 83, 77}
	t.Log(separateDigits(nums))
	nums = []int{7, 1, 3, 9}
	t.Log(separateDigits(nums))
}

/************ 2nd test************/
func maxCount(banned []int, n int, maxSum int) (ans int) {
	dic := make(map[int]bool)
	for _, b := range banned {
		dic[b] = true
	}
	sum := 0
	for i := 1; i <= n && sum+i <= maxSum; i++ {
		if !dic[i] {
			sum += i
			ans++
		}
	}
	return
}
func Test_2nd(t *testing.T) {
	banned := []int{1, 6, 5}
	n, maxSum := 5, 6
	t.Log(maxCount(banned, n, maxSum))
	banned = []int{1, 2, 3, 4, 5, 6, 7}
	n, maxSum = 8, 1
	t.Log(maxCount(banned, n, maxSum))
	banned = []int{11}
	n, maxSum = 7, 50
	t.Log(maxCount(banned, n, maxSum))
}

/************ 3rd test************/
func maximizeWin(prizePositions []int, k int) (ans int) {
	pre := make([]int, len(prizePositions)+1)
	left := 0
	for right, p := range prizePositions {
		for p-prizePositions[left] > k {
			left++
		}
		ans = max(ans, right-left+1+pre[left])
		pre[right+1] = max(pre[right], right-left+1)
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Test_3rd(t *testing.T) {
	prizePositions, k := []int{1, 1, 2, 2, 3, 3, 5}, 2
	t.Log(maximizeWin(prizePositions, k))
	prizePositions, k = []int{1, 2, 3, 4}, 0
	t.Log(maximizeWin(prizePositions, k))
}

/************ 4th test************/

func Test_4th(t *testing.T) {

}
