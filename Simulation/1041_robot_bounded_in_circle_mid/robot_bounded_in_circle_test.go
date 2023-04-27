package robotboundedincircle

import "testing"

func isRobotBounded(instructions string) bool {
	//机器人回到了原点，或者机器人的方向与初始方向不同，那么机器人一定会进入循环
	dir := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} //N,E,S,W
	d := 0
	x, y := 0, 0
	for _, cmd := range instructions {
		switch cmd {
		case 'G':
			x += dir[d][0]
			y += dir[d][1]
		case 'L':
			d = (d + 3) % 4
		case 'R':
			d = (d + 1) % 4
		}
	}
	return (x == 0 && y == 0) || d != 0
}
func Test_robot_bounded_in_circle(t *testing.T) {
	tests := []struct {
		instructions string
		wanna        bool
	}{
		{"GGLLGG", true},
		{"GG", false},
		{"GL", true},
	}
	for _, tt := range tests {
		t.Log(isRobotBounded(tt.instructions) == tt.wanna)
	}
}
