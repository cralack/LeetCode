package splitarrayintoconsecutivesubsequences

import "testing"

func isPossible(nums []int) bool {
	freq := make(map[int]int, 0)
	need := make(map[int]int, 0)
	for _, v := range nums {
		freq[v]++ // 统计 nums 中元素的频率
	}
	for _, v := range nums {
		if freq[v] == 0 { // 已经被用到其他子序列中
			continue
		} // 先判断 v 是否能接到其他子序列后面
		if _, ok := need[v]; ok && need[v] > 0 { // v 可以接到之前的某个序列后面
			freq[v]--
			need[v]--   // 对 v 的需求减一
			need[v+1]++ // 对 v + 1 的需求加一
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 {
			// 将 v 作为开头，新建一个长度为 3 的子序列 [v,v+1,v+2]
			freq[v]--
			freq[v+1]--
			freq[v+2]--
			need[v+3]++ // 对 v + 3 的需求加一
		} else { // 两种情况都不符合，则无法分配
			return false
		}
	}
	return true
}

func Get_Subseq(nums []int) (res [][]int) {
	freq := make(map[int]int, 0)
	need := make(map[int][][]int, 0)

	for _, v := range nums {
		freq[v]++
	}

	for _, v := range nums {
		if freq[v] == 0 {
			continue
		}

		if _, ok := need[v]; ok && len(need[v]) > 0 {
			freq[v]--
			seq := need[v][len(need[v])-1]
			need[v] = need[v][:len(need[v])-1]
			seq = append(seq, v)
			need[v+1] = append(need[v+1], seq)
		} else if freq[v] > 0 && freq[v+1] > 0 && freq[v+2] > 0 {
			freq[v]--
			freq[v+1]--
			freq[v+2]--
			seq := []int{v, v + 1, v + 2}
			need[v+3] = append(need[v+3], seq)
		} else {
			return res
		}
	}
	for _, seqs := range need {
		res = append(res, seqs...)
	}
	return
}
func Test_split_array_into_consecutive_subsequences(t *testing.T) {
	arr := []int{1, 2, 3, 3, 4, 5}
	t.Log(isPossible(arr))
	arr = []int{1, 2, 3, 3, 4, 4, 5, 5}
	t.Log(isPossible(arr))
	arr = []int{1, 2, 3, 4, 4, 5}
	t.Log(isPossible(arr))
	arr = []int{1, 2, 3, 3, 4, 4, 5, 5}
	t.Log(Get_Subseq(arr))
}
