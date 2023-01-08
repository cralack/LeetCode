package weekly_contest

import (
	"container/heap"
	"testing"
)

/************ 1st test************/
func maximumCount(nums []int) int {
	pos, neg := 0, 0
	for _, num := range nums {
		if num < 0 {
			neg++
		} else if num > 0 {
			pos++
		}
	}
	if pos > neg {
		return pos
	}
	return neg
}
func Test_1st(t *testing.T) {
	nums := []int{-2, -1, -1, 1, 2, 3}
	t.Log(maximumCount(nums))
	nums = []int{-3, -2, -1, 0, 0, 1, 2}
	t.Log(maximumCount(nums))
	nums = []int{5, 20, 66, 1314}
	t.Log(maximumCount(nums))
}

/************ 2nd test************/
func maxKelements(nums []int, k int) int64 {
	que := &IntHeap{}
	ans := 0
	for _, n := range nums {
		heap.Push(que, n)
	}
	for i := 0; i < k; i++ {
		cur := heap.Pop(que).(int)
		ans += cur
		cur = (cur + 2) / 3
		heap.Push(que, cur)
	}
	return int64(ans)
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func Test_2nd(t *testing.T) {
	nums, k := []int{10, 10, 10, 10, 10}, 5
	t.Log(maxKelements(nums, k))
	nums, k = []int{1, 10, 3, 3, 3}, 3
	t.Log(maxKelements(nums, k))
}

/************ 3rd test************/
func isItPossible(word1 string, word2 string) bool {
	// 统计每种字符出现次数
	cnt1 := make(map[rune]int, 0)
	for _, c := range word1 {
		cnt1[c]++
	}
	cnt2 := make(map[rune]int, 0)
	for _, c := range word2 {
		cnt2[c]++
	}

	// 枚举交换哪两种字符
	for i := range cnt1 {
		for j := range cnt2 {
			// 交换字符
			cnt1[i]--
			if cnt1[i] == 0 {
				delete(cnt1, i)
			}
			cnt2[i]++
			cnt2[j]--
			if cnt2[j] == 0 {
				delete(cnt2, j)
			}
			cnt1[j]++

			if len(cnt1) == len(cnt2) {
				return true
			}

			cnt1[i]++
			cnt2[i]--
			if cnt2[i] == 0 {
				delete(cnt2, i)
			}
			cnt2[j]++
			cnt1[j]--
			if cnt1[j] == 0 {
				delete(cnt1, j)
			}
		}

	}
	return false
}
func Test_3rd(t *testing.T) {
	word1 := "ac"
	word2 := "b"
	t.Log(isItPossible(word1, word2))
	word1 = "abcc"
	word2 = "aab"
	t.Log(isItPossible(word1, word2))
	word1 = "abcde"
	word2 = "fghij"
	t.Log(isItPossible(word1, word2))
	word1 = "ab"
	word2 = "abcc"
	t.Log(isItPossible(word1, word2))
}

/************ 4th test************/

func Test_4th(t *testing.T) {

}
