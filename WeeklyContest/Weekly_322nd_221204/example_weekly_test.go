package weekly_contest

import (
	"strings"
	"testing"
)

/************ 1st test************/
func isCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")
	first, last := words[0][0], words[0][len(words[0])-1]
	for i := 1; i < len(words); i++ {
		word := words[i]
		n := len(word)
		curFir := word[0]
		curLa := word[n-1]
		if last != curFir {
			return false
		}
		last = curLa
	}
	return first == last
}
func Test_1st(t *testing.T) {
	sentence := "leetcode exercises sound delightful"
	t.Log(isCircularSentence(sentence))
	sentence = "eetcode"
	t.Log(isCircularSentence(sentence))
	sentence = "Leetcode is cool"
	t.Log(isCircularSentence(sentence))
}

/************ 2nd test************/
func dividePlayers(skill []int) (ans int64) {
	sum, n := 0, len(skill)/2
	cnt := make(map[int]int, 0)
	for _, val := range skill {
		sum += val
		cnt[val]++
	}
	if sum%n != 0 {
		return -1
	}
	target := sum / n
	for _, val := range skill {
		if c := cnt[val]; c > 0 {
			ans += int64(val) * int64(target-val)
			cnt[val]--
			cnt[target-val]--
			if cnt[target-val] < 0 {
				return -1
			}
		}
	}
	return
}
func Test_2nd(t *testing.T) {
	skill := []int{3, 2, 5, 1, 3, 4}
	t.Log(dividePlayers(skill))
	skill = []int{3, 4}
	t.Log(dividePlayers(skill))
	skill = []int{1, 1, 2, 3}
	t.Log(dividePlayers(skill))
	skill = []int{2, 1, 5, 2}
	t.Log(dividePlayers(skill))
}

/************ 3rd test************/
func minScore(n int, roads [][]int) (ans int) {
	type edge struct {
		to  int
		dis int
	}
	graph := make([][]edge, n+1)
	vis := make([]bool, n+1)
	for _, road := range roads {
		f, t, d := road[0], road[1], road[2]
		graph[f] = append(graph[f], edge{t, d})
		graph[t] = append(graph[t], edge{f, d})
	}
	ans = 1e4 + 10
	que := []int{1}
	vis[1] = true
	for len(que) > 0 {
		cur := que[0]
		que = que[1:]
		for _, next := range graph[cur] {
			if next.dis < ans {
				ans = next.dis
			}
			if vis[next.to] {
				continue
			}
			vis[next.to] = true
			que = append(que, next.to)
		}
	}

	return
}

func Test_3rd(t *testing.T) {
	roads, n := [][]int{{1, 2, 9}, {2, 3, 6}, {2, 4, 5}, {1, 4, 7}}, 4
	t.Log(minScore(n, roads))
	roads, n = [][]int{{1, 2, 2}, {1, 3, 4}, {3, 4, 7}}, 4
	t.Log(minScore(n, roads))
}

/************ 4th test************/
func magnificentSets(n int, edges [][]int) (ans int) {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		from, to := edge[0], edge[1]
		graph[from] = append(graph[from], to)
		graph[to] = append(graph[to], from)
	}

	bfs := func(start, base int) (mx int) {
		group := make([]int, n+1)
		group[start] = base
		type pair struct{ idx, wei int }
		que := []pair{{start, base}}
		for len(que) > 0 {
			cur := que[0]
			que = que[1:]
			idx, wei := cur.idx, cur.wei
			mx = max(mx, wei)
			for _, next := range graph[idx] {
				if group[next] == 0 {
					group[next] = wei + 1
					que = append(que, pair{next, wei + 1})
				} else if abs(group[next]-group[idx]) != 1 {
					return 0
				}
			}
		}
		return
	}

	vis := make([]bool, n+1)
	for i, b := range vis {
		if b || i == 0 {
			continue
		}
		base := ans + 1 // 接着上个连通分量
		var dfs func(int)
		dfs = func(cur int) {
			ans = max(ans, bfs(cur, base))
			vis[cur] = true
			for _, y := range graph[cur] {
				if !vis[y] {
					dfs(y)
				}
			}
		}
		dfs(i)
		if ans < base { // 没有变大，说明无法找到合法的分组
			return -1
		}
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
func Test_4th(t *testing.T) {
	edges, n := [][]int{{1, 2}, {1, 4}, {1, 5}, {2, 6}, {2, 3}, {4, 6}}, 6
	t.Log(magnificentSets(n, edges))
	edges, n = [][]int{{1, 2}, {2, 3}, {3, 1}}, 3
	t.Log(magnificentSets(n, edges))
}
