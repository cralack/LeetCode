package kokoeatingbananas

import (
	"testing"
)

func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, 0
	//init lo,hi
	for _, pile := range piles {
		if hi < pile {
			hi = pile
		}
	}

	for lo < hi { //O(logn)
		mi := lo + (hi-lo)>>1
		if possible(piles, mi, h) { //k值还不够小
			hi = mi
		} else {
			lo = mi + 1
		}
	}
	return lo //O(n+nlogn)
}
func possible(piles []int, k, h int) bool {
	time := 0 //O(n)
	for _, pile := range piles {
		time += (pile-1)/k + 1 //向上取整
		if time > h {
			return false
		}
	}
	return time <= h
}
func Test_koko_eating_bananas(t *testing.T) {
	piles, h := []int{3, 6, 7, 11}, 8
	t.Log(minEatingSpeed(piles, h))
	piles, h = []int{30, 11, 23, 4, 20}, 5
	t.Log(minEatingSpeed(piles, h))
	piles, h = []int{30, 11, 23, 4, 20}, 6
	t.Log(minEatingSpeed(piles, h))
	piles, h = []int{312884470}, 312884469
	t.Log(minEatingSpeed(piles, h))
}
