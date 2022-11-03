package weekly_contest

import (
	. "LeetCode/BinaryTree/BinTree"
	"testing"
)

/************ 1st test************/
func smallestEvenMultiple(n int) int {
	if n%2 == 1 {
		return n * 2
	}
	return n
}

func Test_1st(t *testing.T) {
	n := 5
	t.Log(smallestEvenMultiple(n))
	n = 6
	t.Log(smallestEvenMultiple(n))
}

/************ 2nd test************/
func longestContinuousSubstring(s string) int {
	maxLen := 1
	cnt := 1
	for i := 1; i < len(s); i++ {
		pre := s[i-1]
		cur := s[i]
		if cur-pre == 1 {
			cnt++
			if maxLen < cnt {
				maxLen = cnt
			}
		} else {
			cnt = 1
		}
	}
	return maxLen
}
func Test_2nd(t *testing.T) {
	s := "abacaba"
	t.Log(longestContinuousSubstring(s))
	s = "abcde"
	t.Log(longestContinuousSubstring(s))
	s = "labccd"
	t.Log(longestContinuousSubstring(s))
	s = "z"
	t.Log(longestContinuousSubstring(s))
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
func reverseOddLevels(root *TreeNode) *TreeNode {
	mat := [][]int{}
	var dfs1 func(*TreeNode, int)
	dfs1 = func(cur *TreeNode, dep int) {
		if cur == nil {
			return
		}
		if len(mat) == dep {
			mat = append(mat, []int{})
		}
		mat[dep] = append(mat[dep], cur.Val)
		dfs1(cur.Left, dep+1)
		dfs1(cur.Right, dep+1)
	}

	var dfs2 func(*TreeNode, int)
	dfs2 = func(cur *TreeNode, dep int) {
		if cur == nil {
			return
		}
		if dep%2 == 1 {
			cur.Val = back(mat[dep])
			mat[dep] = popback(mat[dep])
		}
		dfs2(cur.Left, dep+1)
		dfs2(cur.Right, dep+1)
	}

	dfs1(root, 0)
	dfs2(root, 0)
	return root
}
func back(arr []int) int {
	return arr[len(arr)-1]
}
func popback(arr []int) []int {
	arr = arr[:len(arr)-1]
	return arr
}

func Test_3rd(t *testing.T) {
	//t.Log(pow(4))
	root := Init([]int{2, 3, 5, 8, 13, 21, 34})
	root = reverseOddLevels(root)
	root.Show()
	root = Init([]int{0, 1, 2, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2})
	root = reverseOddLevels(root)
	root.Show()
}

/************ 4th test************/
func sumPrefixScores(words []string) []int {
	//define trie node
	type node struct {
		son   [26]*node
		idx   []int
		score int
	}

	//build trie
	root := &node{}
	for i, word := range words {
		cur := root
		for _, c := range word {
			c -= 'a'
			if cur.son[c] == nil {
				cur.son[c] = &node{}
			}
			cur = cur.son[c]
			cur.score++ // 更新所有前缀的分数
		}
		cur.idx = append(cur.idx, i)
	}

	//
	ans := make([]int, len(words))
	var dfs func(*node, int)
	dfs = func(node *node, sum int) {
		sum += node.score
		for _, i := range node.idx {
			ans[i] = sum
		}
		for _, child := range node.son {
			if child != nil {
				dfs(child, sum)
			}
		}
	}

	dfs(root, 0)
	return ans
}
func Test_4th(t *testing.T) {
	words := []string{"abc", "ab", "bc", "b"}
	t.Log(sumPrefixScores(words))
	words = []string{"abcd"}
	t.Log(sumPrefixScores(words))
}
