package slidingpuzzle

import (
	"strconv"
	"strings"
	"testing"
)

func slidingPuzzle(board [][]int) int {
	const tar = "123450"
	sb := &strings.Builder{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			sb.WriteString(strconv.Itoa(board[i][j]))
		}
	}
	start := sb.String()
	que1 := make(map[string]bool, 0)
	que1[start] = true
	que2 := make(map[string]bool, 0)
	visited := make(map[string]bool, 0)
	que2[tar] = true
	step := 0

	for len(que1) > 0 && len(que2) > 0 {
		if len(que1) > len(que2) { // 优先扩展短队列
			que1, que2 = que2, que1
		} // tmp保存接下来要扩展的节点
		tmp := make(map[string]bool, 0)

		for cur := range que1 {
			if visited[cur] {
				continue
			}
			if que2[cur] { // 两边有交集则直接返回答案
				return step
			}
			visited[cur] = true
			// 对cur节点进行扩展
			next := getNeighbor(cur)
			for _, neighbor := range next {
				if !visited[neighbor] {
					tmp[neighbor] = true
				}
			}
		}
		que1 = que2
		que2 = tmp
		step++
	}
	return -1
}
func getNeighbor(start string) (res []string) {
	neighbor := [][]int{
		{1, 3},
		{0, 4, 2},
		{1, 5},
		{0, 4},
		{3, 1, 5},
		{4, 2}}
	for i, c := range start {
		if c == '0' {
			arr := neighbor[i]
			for _, j := range arr {
				tmp := []byte(start)
				tmp[i], tmp[j] = start[j], start[i]
				res = append(res, string(tmp))
			}
		}
	}
	return res
}
func Test_sliding_puzzle(t *testing.T) {
	board := [][]int{{1, 2, 3}, {4, 0, 5}}
	t.Log(slidingPuzzle(board))
	board = [][]int{{1, 2, 3}, {5, 4, 0}}
	t.Log(slidingPuzzle(board))
	board = [][]int{{4, 1, 2}, {5, 0, 3}}
	t.Log(slidingPuzzle(board))
}
