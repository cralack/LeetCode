package shortestsupersequencelcci

import "testing"

func shortestSeq(big []int, small []int) []int {
	cnt := 0
	idx := 0
	valid := make(map[int]int, 0)
	for _, v := range small {
		valid[v] = 0
	}
	left, right := 0, 0 //slide window idx
	minLen, n := len(big)+1, len(big)
	for left < n && right < n {
		if _, ok := valid[big[right]]; ok { //命中small
			valid[big[right]]++
			if valid[big[right]] == 1 { //第一次出现
				cnt++
			}
		}

		for cnt >= len(small) {
			if _, ok := valid[big[left]]; ok {
				valid[big[left]]--
				if valid[big[left]] == 0 {
					if minLen > right-left {
						minLen = right - left
						idx = left
					}
					cnt--
				}
			}
			left++
		}
		right++
	}

	if minLen < n+1 {
		return []int{idx, idx + minLen}
	} else {
		return []int{}
	}
}
func Test_shortest_supersequence_lcci(t *testing.T) {
	big := []int{7, 5, 9, 0, 2, 1, 3, 5, 7, 9, 1, 1, 5, 8, 8, 9, 7}
	small := []int{1, 5, 9}
	t.Log(shortestSeq(big, small))
	big = []int{1, 2, 3}
	small = []int{4}
	t.Log(shortestSeq(big, small))
}
