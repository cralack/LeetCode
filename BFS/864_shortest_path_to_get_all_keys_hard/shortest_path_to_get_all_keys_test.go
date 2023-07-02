package shortestpathtogetallkeys

import "testing"

func shortestPathAllKeys(grid []string) (ans int) {
	m, n := len(grid), len(grid[0])
	// 统计钥匙数目,起点坐标
	var keyCnt, si, sj int
	for i, row := range grid {
		for j, c := range row {
			if c >= 'a' && c <= 'z' {
				// 累加钥匙数量
				keyCnt++
			} else if c == '@' {
				// 起点
				si, sj = i, j
			}
		}
	}
	// state 表示当前拥有的钥匙的状态，
	// 即 state 的第 i 位为 1 表示当前拥有第 i 把钥匙
	// 0110 代表获得了第b,c把钥匙
	type tuple struct{ i, j, state int }
	start := tuple{si, sj, 0}
	que := []tuple{start}
	vis := map[tuple]bool{start: true}
	dirs := []int{-1, 0, 1, 0, -1}
	for len(que) > 0 {
		for t := len(que); t > 0; t-- {
			cur := que[0]
			que = que[1:]
			i, j, state := cur.i, cur.j, cur.state
			// 找到所有钥匙，返回当前步数
			// 假如k=3代表需要拿到3把钥匙
			// 1 << 3 = 1000 再减1为 111
			if 1<<keyCnt-1 == state {
				return ans
			}
			// 往四个方向搜索
			for d := 0; d < 4; d++ {
				x, y := i+dirs[d], j+dirs[d+1]
				// 在边界范围内
				if x >= 0 && x < m && y >= 0 && y < n {
					c := grid[x][y]
					// 是墙，或者是锁，但此时没有对应的钥匙，无法通过
					if c == '#' || (c >= 'A' && c <= 'Z' && (state>>(c-'A')&1 == 0)) {
						continue
					}
					nxt := state
					// 是钥匙，更新状态
					if c >= 'a' && c <= 'z' {
						nxt |= 1 << (c - 'a')
					}
					// 此状态未访问过，入队
					if !vis[tuple{x, y, nxt}] {
						vis[tuple{x, y, nxt}] = true
						que = append(que, tuple{x, y, nxt})
					}
				}
			}
		}
		// 步数加一
		ans++
	}
	return -1
}
func Test_shortest_path_to_get_all_keys(t *testing.T) {
	grid := []string{"@.a.#", "###.#", "b.A.B"}
	t.Log(shortestPathAllKeys(grid))
	grid = []string{"@..aA", "..B#.", "....b"}
	t.Log(shortestPathAllKeys(grid))
	grid = []string{"@Aa"}
	t.Log(shortestPathAllKeys(grid))
}
