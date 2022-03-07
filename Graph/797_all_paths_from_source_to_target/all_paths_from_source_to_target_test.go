package allpathsfromsourcetotarget

import (
	"testing"
)

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}
	path := []int{}
	var traverse func(s int)
	traverse = func(s int) {
		path = append(path, s)
		size := len(graph)
		if s == size-1 {
			res = append(res, append([]int(nil), path...))
			path = path[:len(path)-1]
			return
		}

		for _, ver := range graph[s] {
			traverse(ver)
		}

		path = path[:len(path)-1]
	}

	traverse(0)
	return res
}
func Test_all_paths_from_source_to_target(t *testing.T) {
	graph := [][]int{
		/*0:*/ {3, 1},
		/*1:*/ {4, 6, 7, 2, 5},
		/*2:*/ {4, 6, 3},
		/*3:*/ {6, 4},
		/*4:*/ {7, 6, 5},
		/*5:*/ {6},
		/*6:*/ {7},
		/*7:*/ {}}
	ans := allPathsSourceTarget(graph)
	t.Log(ans)
}
