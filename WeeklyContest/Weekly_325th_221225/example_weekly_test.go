package weekly_contest

import (
	"sort"
	"testing"
)

/************ 1st test************/
func closetTarget(words []string, target string, startIndex int) (ans int) {
	n := len(words)
	ans = n
	for i, w := range words {
		if w == target {
			//用正向反向的距离更新ans
			ans = min(ans, min((startIndex+n-i)%n, (i-startIndex+n)%n))
		}
	}
	if ans == n {
		return -1
	}
	return
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Test_1st(t *testing.T) {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	target := "hello"
	startIndex := 1
	t.Log(closetTarget(words, target, startIndex))
	words = []string{"a", "b", "leetcode"}
	target = "leetcode"
	startIndex = 0
	t.Log(closetTarget(words, target, startIndex))
	words = []string{"i", "eat", "leetcode"}
	target = "ate"
	startIndex = 0
	t.Log(closetTarget(words, target, startIndex))
	words = []string{"hsdqinnoha", "mqhskgeqzr", "zemkwvqrww", "zemkwvqrww", "daljcrktje", "fghofclnwp", "djwdworyka", "cxfpybanhd", "fghofclnwp", "fghofclnwp"}
	target = "zemkwvqrww"
	startIndex = 8
	t.Log(closetTarget(words, target, startIndex))
}

/************ 2nd test************/
func takeCharacters(s string, k int) (ans int) {
	n := len(s)
	c, j := [3]int{}, n
	//从右往左取到第j个字符
	for c[0] < k || c[1] < k || c[2] < k {
		if j == 0 {
			return -1
		}
		j--
		c[s[j]-'a']++
	}
	// 左侧没有取字符
	ans = n - j
	//枚举从左侧取到第i个字符
	for i := 0; i < n && j < n; i++ {
		c[s[i]-'a']++
		// 维护 j 的最大下标
		for j < n && c[s[j]-'a'] > k {
			c[s[j]-'a']--
			j++
		}
		ans = min(ans, i+n-j+1)
	}
	return
}

func Test_2nd(t *testing.T) {
	s, k := "aabaaaacaabc", 2
	t.Log(takeCharacters(s, k))
	s, k = "a", 1
	t.Log(takeCharacters(s, k))
	s, k = "a", 0
	t.Log(takeCharacters(s, k))
	s, k = "abc", 1
	t.Log(takeCharacters(s, k))
	s, k = "ccbabcc", 1
	t.Log(takeCharacters(s, k))
}

/************ 3rd test************/
func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	left, right := 0, price[len(price)-1]
	for left < right {
		diff := left + (right-left)>>1
		cnt := 1
		pre := price[0]
		for _, cur := range price[1:] {
			if cur-pre >= diff {
				cnt++
				pre = cur
			}
		}
		//不满足条件 diff取大了
		if cnt < k {
			right = diff
		} else { //diff取小了
			left = diff + 1
		}
	} //二分出不满足条件的最小值
	return left - 1
}
func Test_3rd(t *testing.T) {
	price, k := []int{13, 5, 1, 8, 21, 2}, 3
	t.Log(maximumTastiness(price, k))
	price, k = []int{1, 3, 1}, 2
	t.Log(maximumTastiness(price, k))
	price, k = []int{7, 7, 7, 7}, 2
	t.Log(maximumTastiness(price, k))

}

/************ 4th test************/

func Test_4th(t *testing.T) {

}
