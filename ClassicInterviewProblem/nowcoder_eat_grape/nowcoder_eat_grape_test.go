package nowcodereatgrape

import (
	"sort"
	"testing"
)

/*
链接：https://www.nowcoder.com/questionTerminal/14c0359fb77a48319f0122ec175c9ada
来源：牛客网

有三种葡萄，每种分别有a,b,c颗。有三个人，
第一个人只吃第1,2种葡萄，
第二个人只吃第2,3种葡萄，
第三个人只吃第1,3种葡萄。
适当安排三个人使得吃完所有的葡萄,
并且且三个人中吃的最多的那个人吃得尽量少。
*/
func solution(arr []int) int {
	sort.Ints(arr)
	a, b, c := arr[0], arr[1], arr[2]
	sum := a + b + c
	if a+b > c { // 能够构成三角形，可完全平分
		return (sum + 2) / 3
	}
	if 2*(a+b) < c { // 不能构成三角形，平分最长边的情况
		return (c + 1) / 2
	}
	return (sum + 2) / 3 // 不能构成三角形，但依然可以完全平分的情况
}
func Test_nowcoder_eat_grape(t *testing.T) {
	grapes := []int{1, 2, 3}
	t.Log(solution(grapes))
	grapes = []int{1, 2, 6}
	t.Log(solution(grapes))
	grapes = []int{12, 13, 11}
	t.Log(solution(grapes))

}
