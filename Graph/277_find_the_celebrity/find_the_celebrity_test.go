package findthecelebrity

import (
	"testing"
)

func findCelebrity(n int) int {
	if n == 1 {
		return 0
	}
	q := []int{}
	for i := 0; i < n; i++ {
		q = append(q, i)
	}
	for len(q) >= 2 {
		cand := q[0]
		q = q[1:]
		other := q[0]
		q = q[1:]
		if knows(cand, other) || !knows(other, cand) {
			q = append(q, other)
		} else {
			q = append(q, cand)
		}
	}
	cand := q[0]
	for other := 0; other < n; other++ {
		if other == cand {
			continue
		} // 保证其他人都认识 cand，且 cand 不认识任何其他人
		if !knows(other, cand) || knows(cand, other) {
			return -1
		}
	}
	return cand
}
func findCelebrity_better(n int) int {
	cand := 0
	for other := 1; other < n; other++ {
		if !knows(other, cand) || knows(cand, other) {
			cand = other
		}
	}
	for other := 0; other < n; other++ {
		if cand == other {
			continue
		}
		if !knows(other, cand) || knows(cand, other) {
			return -1
		}
	}
	return cand
}
func knows(i, j int) bool {
	graph := [][]int{
		{1, 1, 0},
		{0, 1, 0},
		{1, 1, 1}}
	tmp := graph[i][j]
	return tmp == 1
}
func Test_find_the_celebrity(t *testing.T) {
	t.Log(findCelebrity(3))
	t.Log(findCelebrity_better(3))
}
