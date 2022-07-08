package sequencereconstructionmid

import "testing"

func sequenceReconstruction(nums []int, sequences [][]int) bool {
	n := len(nums)
	inDegree := make([]int, n+1) //[1，n]
	graph := make([][]int, n+1)
	for _, seq := range sequences { //建立有向图
		for i := 0; i < len(seq)-1; i++ {
			from, to := seq[i], seq[i+1]
			graph[from] = append(graph[from], to)
			inDegree[to]++
		}
	}

	queue := []int{}
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 { //0出度即起点
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		if len(queue) > 1 { //序列构建不唯一则false
			return false
		}
		cur := queue[0] //pop
		queue = queue[1:]
		for _, nextNode := range graph[cur] {
			if inDegree[nextNode]--; inDegree[nextNode] == 0 {
				queue = append(queue, nextNode)
			}
		}
	}

	return true
}
func Test_sequence_reconstruction(t *testing.T) {
	nums := []int{1, 2, 3}
	sequences := [][]int{{1, 2}, {1, 3}}
	t.Log(sequenceReconstruction(nums, sequences))
	nums = []int{1, 2, 3}
	sequences = [][]int{{1, 2}}
	t.Log(sequenceReconstruction(nums, sequences))
	nums = []int{1, 2, 3}
	sequences = [][]int{{1, 2}, {1, 3}, {2, 3}}
	t.Log(sequenceReconstruction(nums, sequences))
}
