package weeklycompetition

import (
	"container/heap"
	"sort"
	"testing"
)

/************ 1st test************/
func mostFrequentEven(nums []int) (ans int) {
	max := 0
	cnt := make(map[int]int, 0)
	for _, n := range nums {
		if n%2 == 0 {
			cnt[n]++
			if max < cnt[n] {
				max = cnt[n]
			}
		}
	}
	if len(cnt) == 0 {
		return -1
	}
	ans = 1e+5
	for k, v := range cnt {
		if v == max && ans > k {
			ans = k
		}
	}
	return
}
func Test_1st(t *testing.T) {
	nums := []int{0, 1, 2, 2, 4, 4, 1}
	t.Log(mostFrequentEven(nums))
	nums = []int{4, 4, 4, 9, 2, 4}
	t.Log(mostFrequentEven(nums))
	nums = []int{29, 47, 21, 41, 13, 37, 25, 7}
	t.Log(mostFrequentEven(nums))
}

/************ 2nd test************/
func partitionString(s string) (ans int) {
	ans = 1
	vis := make(map[rune]bool, 0)
	for _, c := range s {
		if _, has := vis[c]; has {
			ans++
			vis = make(map[rune]bool, 26)
		}
		vis[c] = true
	}
	return
}
func Test_2nd(t *testing.T) {
	s := "abacaba"
	t.Log(partitionString(s))
	s = "ssssss"
	t.Log(partitionString(s))
}

/************ 3rd test************/
func minGroups(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	pq := IntHeap{}
	for _, interval := range intervals {
		if !pq.Empty() && pq[0] < interval[0] {
			heap.Pop(&pq)
		}
		heap.Push(&pq, interval[1])
	}
	return pq.Len()
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Empty() bool        { return h.Len() == 0 }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func Test_3rd(t *testing.T) {
	intervals := [][]int{{5, 10}, {6, 8}, {1, 5}, {2, 3}, {1, 10}}
	t.Log(minGroups(intervals))
	intervals = [][]int{{1, 3}, {5, 6}, {8, 10}, {11, 13}}
	t.Log(minGroups(intervals))
}

/************ 4th test************/
//TLE
// func lengthOfLIS(nums []int, k int) int {
// 	n := len(nums)
// 	dp := make([]int, n)
// 	for i := range dp {
// 		dp[i] = 1
// 	}
// 	for i := 0; i < n; i++ {
// 		for j := 0; j < i; j++ {
// 			if nums[i] > nums[j] && nums[i]-nums[j] <= k {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 	}
// 	res := 0
// 	for i := range dp {
// 		res = max(res, dp[i])
// 	}
// 	return res
// }
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// 线段树解法
func lengthOfLIS(nums []int, k int) (ans int) {
	const mx = 1e5
	t := buildST(mx)
	for _, v := range nums {
		// 此处没有特殊处理v=1及v<=k的情况，因为区间查询时能够兼容
		t.update(v, 1+t.query(v-k, v-1))
	}
	return t[1].max
}

type seg []struct{ l, r, max int }

// buildST 构造线段树
func buildST(n int) (t seg) {
	t = make(seg, n<<2)
	var dfs func(o, l, r int)
	dfs = func(o, l, r int) {
		t[o].l, t[o].r = l, r
		if l != r {
			mid, ol := t.mid(o), o<<1
			dfs(ol, l, mid)
			dfs(ol|1, mid+1, r)
		}
	}
	dfs(1, 1, n)
	return
}

func (t seg) mid(o int) int {
	return (t[o].l + t[o].r) >> 1
}

// 单点更新
func (t seg) update(i, val int) {
	var dfs func(o int)
	dfs = func(o int) {
		if t[o].l == t[o].r {
			t[o].max = val
		} else {
			mid, ol := t.mid(o), o<<1
			if i <= mid {
				dfs(ol) // i位于左子树
			} else {
				dfs(ol | 1) // i位于右子树
			}
			// 更新o节点维护的区间信息
			t[o].max = MaxInt(t[ol].max, t[ol|1].max)
		}
	}
	dfs(1)
}

// 区间查询
func (t seg) query(l, r int) int {
	var dfs func(o int) (ans int)
	dfs = func(o int) (ans int) {
		if l <= t[o].l && t[o].r <= r { // 区间[l, r]包含o节点对应的区间[t[o].l, t[o].r]
			return t[o].max
		} else {
			mid, ol := t.mid(o), o<<1
			if r <= mid {
				return dfs(ol) // 查询左子树
			} else if l > mid {
				return dfs(ol | 1) // 查询右子树
			} else {
				return MaxInt(dfs(ol), dfs(ol|1)) // 查询左右子树
			}
		}
	}
	return dfs(1)
}

func MaxInt(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Test_4th(t *testing.T) {
	nums := []int{4, 2, 1, 4, 3, 4, 5, 8, 15}
	k := 3
	t.Log(lengthOfLIS(nums, k))
	nums = []int{7, 4, 5, 1, 8, 12, 4, 7}
	k = 5
	t.Log(lengthOfLIS(nums, k))
	nums = []int{1, 5}
	k = 1
	t.Log(lengthOfLIS(nums, k))
}
