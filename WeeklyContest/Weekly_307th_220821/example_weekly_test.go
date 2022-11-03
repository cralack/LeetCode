package weekly_contest

import (
	. "LeetCode/BinaryTree/BinTree"
	"container/heap"
	"sort"
	"testing"
)

/************ 1st test************/
func minNumberOfHours(initialEnergy int, initialExperience int,
	energy []int, experience []int) int {
	ans := 0
	n := len(energy)
	for i := 0; i < n; i++ {
		// 精力上无法打败对手，因此最开始需要增加一些精力，使得精力恰好满足要求
		if initialEnergy <= energy[i] {
			delta := energy[i] - initialEnergy + 1
			initialEnergy += delta
			ans += delta
		}
		// 经验上无法打败对手，因此最开始需要增加一些经验，使得经验恰好满足要求
		if initialExperience <= experience[i] {
			delta := experience[i] - initialExperience + 1
			initialExperience += delta
			ans += delta
		}
		// 打败对手后数值变更
		initialEnergy -= energy[i]
		initialExperience += experience[i]
	}
	return ans
}
func Test_1st(t *testing.T) {
	initialEnergy := 5
	initialExperience := 3
	energy := []int{1, 4, 3, 2}
	experience := []int{2, 6, 3, 1}
	t.Log(minNumberOfHours(initialEnergy, initialExperience, energy, experience))
	initialEnergy = 2
	initialExperience = 4
	energy = []int{1}
	experience = []int{3}
	t.Log(minNumberOfHours(initialEnergy, initialExperience, energy, experience))
	initialEnergy = 1
	initialExperience = 1
	energy = []int{1, 1, 1, 1}
	experience = []int{1, 1, 1, 50}
	t.Log(minNumberOfHours(initialEnergy, initialExperience, energy, experience))
}

/************ 2nd test************/
func largestPalindromic(num string) string {
	cnt := make([]int, 10)
	for _, c := range num {
		cnt[(c-'0')]++
	}
	// ans 表示对应的部分中的前一半，ans2 是 ans 的倒序
	var str1, str2 []byte
	// 求回文串两边对应的部分
	for i := 9; i >= 0; i-- {
		//已经枚举到了 0，但是之前从来没有加入过别的数字。
		//此时加入 0 将会导致前导 0，因此直接结束。
		if i == 0 && len(str1) == 0 {
			break
		}
		// 在这部分中出现过的数必须出现偶数次
		t := cnt[i] / 2
		for j := 0; j < t; j++ {
			str1 = append(str1, byte(i+'0'))
		}
		cnt[i] -= t * 2
	}
	//reverse
	n := len(str1)
	str2 = make([]byte, n)
	for i := 0; i < n; i++ {
		str2[i] = str1[n-i-1]
	}
	// 看看是否还有剩下的数，可以作为中间单独的一个数字
	for i := 9; i >= 0; i-- {
		if cnt[i] > 0 {
			// 此时 0 无需跳过，因为单独一个 0 是合法的答案
			str1 = append(str1, byte(i+'0'))
			break
		}
	}
	return string(append(str1, str2...))
}
func Test_largest_palindromic_number(t *testing.T) {
	num := "444947137"
	t.Log(largestPalindromic(num))
}

/************ 3rd test************/
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func amountOfTime(root *TreeNode, start int) int {
	var mx = 0
	var arr [][]int

	//求树中最大值mx,确定arr边界
	var dfs1 func(*TreeNode)
	dfs1 = func(cur *TreeNode) {
		if cur == nil {
			return
		}
		mx = max(mx, cur.Val)
		dfs1(cur.Left)
		dfs1(cur.Right)
	}
	//构建无根树(邻接矩阵)
	var dfs2 func(*TreeNode)
	dfs2 = func(cur *TreeNode) {
		if cur.Left != nil {
			arr[cur.Val] = append(arr[cur.Val], cur.Left.Val)
			arr[cur.Left.Val] = append(arr[cur.Left.Val], cur.Val)
			dfs2(cur.Left)
		}
		if cur.Right != nil {
			arr[cur.Val] = append(arr[cur.Val], cur.Right.Val)
			arr[cur.Right.Val] = append(arr[cur.Right.Val], cur.Val)
			dfs2(cur.Right)
		}
	}
	// DFS求深度
	var dfs3 func(int, int) int
	dfs3 = func(cur, comefrom int) int {
		ret := 0
		for _, next := range arr[cur] {
			if next != comefrom {
				ret = max(ret, dfs3(next, cur)+1)
			}
		}
		return ret
	}

	dfs1(root)
	arr = make([][]int, mx+1)
	dfs2(root)
	return dfs3(start, -1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_amount_of_time_for_binary_tree_to_be_infected(t *testing.T) {
	nums := []int{1, 5, 3, -1, 4, 10, 6, 9, 2}
	tar := 3
	root := Init(nums)
	t.Log(amountOfTime(root, tar))
}

/************ 4th test************/
func kSum(nums []int, k int) int64 {
	n := len(nums)
	sum := 0
	for i, x := range nums {
		if x >= 0 {
			sum += x
		} else {
			nums[i] = -x
		}
	}
	sort.Ints(nums)
	h := &hp{{sum, 0}}
	for ; k > 1; k-- {
		p := heap.Pop(h).(pair)
		if p.i < n {
			heap.Push(h, pair{p.sum - nums[p.i], p.i + 1}) // 保留 nums[p.i-1]
			if p.i > 0 {
				heap.Push(h, pair{p.sum - nums[p.i] + nums[p.i-1], p.i + 1}) // 不保留 nums[p.i-1]
			}
		}
	}
	return int64((*h)[0].sum)
}

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int            { return len(h) }
func (h hp) Less(i, j int) bool  { return h[i].sum > h[j].sum }
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

func Test_kSum(t *testing.T) {
	nums := []int{2, 4, -2}
	k := 5
	t.Log(kSum(nums, k))

	nums = []int{1, -2, 3, 4, -10, 12}
	k = 16
	t.Log(kSum(nums, k))
}
