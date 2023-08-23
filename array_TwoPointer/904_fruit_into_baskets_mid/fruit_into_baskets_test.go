package fruitintobaskets

import "testing"

func totalFruit(fruits []int) (ans int) {
	cnt := map[int]int{}
	left := 0
	for right, cur := range fruits {
		cnt[cur]++
		for len(cnt) > 2 {
			pre := fruits[left]
			cnt[pre]--
			if cnt[pre] == 0 {
				delete(cnt, pre)
			}
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_fruit_into_baskets(t *testing.T) {
	fruits := []int{1, 2, 1}
	t.Log(totalFruit(fruits))
	fruits = []int{0, 1, 2, 2}
	t.Log(totalFruit(fruits))
	fruits = []int{1, 2, 3, 2, 2}
	t.Log(totalFruit(fruits))
	fruits = []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	t.Log(totalFruit(fruits))
}
