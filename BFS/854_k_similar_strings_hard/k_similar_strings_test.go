package ksimilarstringshard

import (
	"testing"
)

// var n int
// var t string
// var mp map[string]int

// func kSimilarity_Astar(start string, target string) int {
// 	n := len(start)
// 	t = target
// 	if start == target {
// 		return 0
// 	}
// 	mp = make(map[string]int, 0)
// 	pq := &StringHeap{}
// 	mp[start] = 0
// 	heap.Push(pq, start)
// 	for pq.Len() != 0 {
// 		poll := heap.Pop(pq).(string)
// 		step := mp[poll]
// 		cs := []byte(poll)
// 		idx := 0
// 		for idx < n && cs[idx] == t[idx] {
// 			idx++
// 		}
// 		for i := idx + 1; i < n; i++ {
// 			if cs[i] == t[idx] && cs[i] != t[i] {
// 				cs[idx], cs[i] = cs[i], cs[idx]
// 				nstr := string(cs)
// 				if val, has := mp[nstr]; has && val <= step+1 {
// 					continue
// 				}
// 				if nstr == t {
// 					return step + 1
// 				}
// 				mp[nstr] = step + 1
// 				heap.Push(pq, nstr)
// 				cs[idx], cs[i] = cs[i], cs[idx]
// 			}
// 		}
// 	}
// 	return -1
// }

// func foo(str string) int {
// 	ans := 0
// 	for i := 0; i < n; i++ {
// 		if str[i] != t[i] {
// 			ans += 1
// 		}
// 	}
// 	return ans + 1>>1
// }

// type StringHeap []string

// func (h StringHeap) Len() int { return len(h) }
// func (h StringHeap) Less(i, j int) bool {
// 	v1, v2 := foo(h[i]), foo(h[j])
// 	d1, d2 := mp[h[i]], mp[h[j]]
// 	return (v1+d1)-(v2+d2) > 0
// }
// func (h StringHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

// func (h *StringHeap) Push(x interface{}) {
// 	*h = append(*h, x.(string))
// }

// func (h *StringHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

func kSimilarity(start string, target string) int {
	if start == target {
		return 0
	}
	nextStrs := func(str string) (ans []string) {
		i := 0
		for ; str[i] == target[i]; i++ {
		}
		for j := i + 1; j < len(start); j++ {
			tmp := []byte(str)
			if tmp[j] == target[i] && tmp[j] != target[j] {
				tmp[i], tmp[j] = tmp[j], tmp[i]
				ans = append(ans, string(tmp))
			}
		}
		return
	}

	vis := make(map[string]bool, 0)
	que := make(map[string]bool, 0)
	que[start] = true
	step := 0
	for len(que) > 0 {
		tmp := que
		que = make(map[string]bool)

		for cur := range tmp {
			if tmp[target] {
				return step
			}
			vis[cur] = true
			nexts := nextStrs(cur)
			for _, next := range nexts {
				if !vis[next] {
					que[next] = true
				}
			}
		}
		step++
	}
	return -1
}

func Test_k_similar_strings(t *testing.T) {
	s1 := "ab"
	s2 := "ba"
	t.Log(kSimilarity(s1, s2))
	s1 = "abc"
	s2 = "bca"
	t.Log(kSimilarity(s1, s2))
	s1 = "abccaacceecdeea"
	s2 = "bcaacceeccdeaae"
	t.Log(kSimilarity(s1, s2))
}
