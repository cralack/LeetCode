package minimumelements

import "testing"

func minElements(nums []int, limit int, goal int) (ans int) {
	tar := 0
	for _, n := range nums {
		tar += n
	}
	tar -= goal
	if tar < 0 {
		tar = -tar
	}
	// 向上取整
	return (tar + limit - 1) / limit
}

func Test_minimum_elements_to_add_to_form_a_given_sum(t *testing.T) {
	nums, limit, goal := []int{1, -1, 1}, 3, -4
	t.Log(minElements(nums, limit, goal))
	nums, limit, goal = []int{1, -10, 9, 1}, 100, 0
	t.Log(minElements(nums, limit, goal))
}
