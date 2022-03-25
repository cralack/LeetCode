package countcompletetreenodes

import (
	. "Learning/LeetCode/BinaryTree/BinTree"
	"fmt"
	"math"
	"testing"
)

func CountNodes1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + CountNodes1(root.Left) + CountNodes1(root.Right)
}
func CountNodes2(root *TreeNode) int {
	l, r := root, root
	hL, hR := 0, 0
	for l != nil {
		l = l.Left
		hL++
	}
	for r != nil {
		r = r.Right
		hR++
	}
	if hL == hR {
		return int(math.Pow(2, float64(hL))) - 1
	}
	return 1 + CountNodes2(root.Left) + CountNodes2(root.Right)
}

func Test_count_complete_tree_nodes(t *testing.T) {
	arr := []int{}
	for i := 0; i < 1e5; i++ {
		arr = append(arr, i)
	}
	root := Init(arr)
	t.Log(Serialize(root))
}
func Benchmark_count_complete_tree_nodes(b *testing.B) {
	arr := []int{}
	for i := 0; i < 1e2; i++ {
		arr = append(arr, i)
	}
	root := Init(arr)
	fmt.Println(Serialize(root))
	b.Run("normal bintree", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CountNodes1(root)
		}
		b.StopTimer()
	})
	b.Run("full bintree", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			CountNodes2(root)
		}
		b.StopTimer()
	})
}
