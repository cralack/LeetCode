package LCCUP

import (
	"sort"
	"strconv"
	"testing"
)

/************ 1st test************/
func temperatureTrend(temperatureA []int, temperatureB []int) (ans int) {
	n := len(temperatureA)
	cnt := 0
	for i := 1; i < n; i++ {
		difA := temperatureA[i] - temperatureA[i-1]
		difB := temperatureB[i] - temperatureB[i-1]
		if (difA > 0 && difB > 0) || (difA < 0 && difB < 0) || (difA == 0 && difB == 0) {
			cnt++
		} else {
			cnt = 0
		}
		if ans < cnt {
			ans = cnt
		}
	}
	return
}
func Test_1st(t *testing.T) {
	temperatureA := []int{21, 18, 18, 18, 31}
	temperatureB := []int{34, 32, 16, 16, 17}
	t.Log(temperatureTrend(temperatureA, temperatureB))
	temperatureA = []int{5, 10, 16, -6, 15, 11, 3}
	temperatureB = []int{16, 22, 23, 23, 25, 3, -16}
	t.Log(temperatureTrend(temperatureA, temperatureB))
}

/************ 2nd test************/
func transportationHub(path [][]int) int {
	inDegree := make(map[int][]int, 0)
	outDegree := make(map[int][]int, 0)
	node := make(map[int]bool, 0)

	for _, p := range path {
		from, to := p[0], p[1]
		node[from] = true
		node[to] = true
		inDegree[to] = append(inDegree[to], from)
		outDegree[from] = append(outDegree[from], to)

	}
	n := len(node)
	candidate, cnt := -1, 0
	for can, inDs := range inDegree {
		if cnt < len(inDs) {
			cnt = len(inDs)
			candidate = can
		}
	}
	if _, has := outDegree[candidate]; !has &&
		candidate != -1 && len(inDegree[candidate]) == n-1 {
		return candidate
	}
	return -1
}
func Test_2nd(t *testing.T) {
	path := [][]int{{0, 1}, {0, 3}, {1, 3}, {2, 0}, {2, 3}}
	t.Log(transportationHub(path))
	path = [][]int{{0, 3}, {1, 0}, {1, 3}, {2, 0}, {3, 0}, {3, 2}}
	t.Log(transportationHub(path))
	path = [][]int{{2, 5}, {4, 3}, {2, 3}}
	t.Log(transportationHub(path))
}

/************ 3rd test************/
func ballGame(num int, plate []string) (ans [][]int) {
	n, m := len(plate), len(plate[0])
	starts := [][]int{}
	for i := range plate {
		for j := range plate[i] {
			if plate[i][j] == 'O' {
				starts = append(starts, []int{i, j})
			}
		}
	}

	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	try := func(start []int, dir int) []int {
		step := 0
		curX, curY := start[0], start[1]
		for step < num {
			nextX := curX + dirs[dir][0]
			nextY := curY + dirs[dir][1]
			if n-1 < nextX || nextX < 0 || m-1 < nextY || nextY < 0 {
				break
			}
			if plate[nextX][nextY] == 'E' {
				dir += 3
			} else if plate[nextX][nextY] == 'W' {
				dir += 1
			}
			step++
			dir %= 4
			curX = nextX
			curY = nextY
		}
		if (curX == 0 && curY == 0) || (curX == 0 && curY == m-1) ||
			(curX == n-1 && curY == 0) || (curX == n-1 && curY == m-1) ||
			plate[curX][curY] == 'O' {
			return []int{-1, -1}
		}
		if step == num && (0 < curX && curX < n-1 && 0 < curY && curY < m-1) {
			return []int{-1, -1}
		}
		return []int{curX, curY}
	}
	set := make(map[string]bool, 0)
	for _, s := range starts {
		for i := 0; i < 4; i++ {
			start := try(s, i)
			if start[0] == -1 {
				continue
			} else {
				if _, has := set[strconv.Itoa(start[0])+","+
					strconv.Itoa(start[1])]; !has {
					set[strconv.Itoa(start[0])+","+strconv.Itoa(start[1])] = true
					ans = append(ans, start)
				}
			}
		}
	}
	sort.Slice(ans, func(i, j int) bool {
		return ans[i][0] < ans[j][0] ||
			ans[i][0] == ans[j][0] && ans[i][1] < ans[j][1]
	})
	return
}

func Test_3rd(t *testing.T) {
	plate, num := []string{"..E.", ".EOW", "..W."}, 4
	t.Log(ballGame(num, plate))
	plate, num = []string{".....", "..E..", ".WO..", "....."}, 5
	t.Log(ballGame(num, plate))
	plate, num = []string{".....", "....O", "....O", "....."}, 3
	t.Log(ballGame(num, plate))
	plate, num = []string{".....", "..O..", "....."}, 1
	t.Log(ballGame(num, plate))
	plate, num = []string{
		"E...W..WW",
		".E...O...",
		"...WO...W",
		"..OWW.O..",
		".W.WO.W.E",
		"O..O.W...",
		".OO...W..",
		"..EW.WEE."}, 41
	t.Log(ballGame(num, plate))
}

/************ 4th test************/

func Test_4th(t *testing.T) {

}

/************ 5th test************/

func Test_5th(t *testing.T) {

}
