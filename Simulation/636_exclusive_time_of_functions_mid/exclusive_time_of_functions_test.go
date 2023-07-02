package exclusivetimeoffunctionsmid

import (
	"strconv"
	"strings"
	"testing"
)

type App struct {
	Id        int
	TimeStamp int
}

func exclusiveTime(n int, logs []string) []int {
	ans := make([]int, n)
	stack := []App{} // 存入app的id以及操作时间
	for _, log := range logs {
		ana := strings.Split(log, ":")
		curId, _ := strconv.Atoi(ana[0])
		timeStamp, _ := strconv.Atoi(ana[2])

		if ana[1][0] == 's' {
			if len(stack) > 0 {
				ans[stack[len(stack)-1].Id] +=
					timeStamp - stack[len(stack)-1].TimeStamp
				// always rewrite by pop↓
				// stack[len(stack)-1].TimeStamp = timeStamp
			}
			stack = append(stack, App{curId, timeStamp})
		} else {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[cur.Id] += timeStamp - cur.TimeStamp + 1
			if len(stack) > 0 {
				stack[len(stack)-1].TimeStamp = timeStamp + 1
			}
		}
	}
	return ans
}
func Test_exclusive_time_of_functions(t *testing.T) {
	logs, n := []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}, 2
	t.Log(exclusiveTime(n, logs))
	logs, n = []string{"0:start:0", "0:start:2", "0:end:5", "0:start:6",
		"0:end:6", "0:end:7"}, 1
	t.Log(exclusiveTime(n, logs))
	logs, n = []string{"0:start:0", "0:start:2", "0:end:5", "1:start:6",
		"1:end:6", "0:end:7"}, 2
	t.Log(exclusiveTime(n, logs))
	logs, n = []string{"0:start:0", "0:start:2", "0:end:5", "1:start:7",
		"1:end:7", "0:end:8"}, 2
	t.Log(exclusiveTime(n, logs))
	logs, n = []string{"0:start:0", "0:end:0"}, 1
	t.Log(exclusiveTime(n, logs))
}
