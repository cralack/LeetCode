package walking_robot_simulation

import (
	"testing"
)

func robotSim(commands []int, obstacles [][]int) (ans int) {
	var (
		dir        = [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
		d          = 0
		curX, curY = 0, 0
	)
	type pair struct {
		x, y int
	}

	obs := make(map[pair]struct{})
	for _, ob := range obstacles {
		obs[pair{ob[0], ob[1]}] = struct{}{}
	}
	for _, c := range commands {
		switch c {
		case -1:
			d = (d + 1) % 4
		case -2:
			d = (d + 3) % 4
		default:
			for i := 0; i < c; i++ {
				if _, has := obs[pair{curX + dir[d][0], curY + dir[d][1]}]; has {
					break
				}
				curX, curY = curX+dir[d][0], curY+dir[d][1]
				if ans < curX*curX+curY*curY {
					ans = curX*curX + curY*curY
				}
			}
		}
	}
	return
}

func Test_walking_robot_simulation(t *testing.T) {
	tests := []struct {
		commands  []int
		obstacles [][]int
	}{
		{[]int{4, -1, 3}, [][]int{}},
		{[]int{4, -1, 4, -2, 4}, [][]int{{2, 4}}},
		{[]int{6, -1, -1, 6}, [][]int{}},
	}
	for _, tt := range tests {
		t.Log(robotSim(tt.commands, tt.obstacles))
	}
}
