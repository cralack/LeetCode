package numberofrecentcalls

import "testing"

type RecentCounter struct {
	queue []int
}

func Constructor() (_ RecentCounter) {
	return
}

func (p *RecentCounter) Ping(t int) int {
	p.queue = append(p.queue, t)
	for p.queue[0] < t-3000 {
		p.queue = p.queue[1:]
	}
	return len(p.queue)
}

func Test_number_of_recent_calls(t *testing.T) {
	c1 := []string{"RecentCounter", "ping", "ping", "ping", "ping"}
	c2 := []int{-1, 1, 100, 3001, 3002}
	var rec RecentCounter
	for i, c := range c1 {
		switch c {
		case "RecentCounter":
			rec = Constructor()
		case "ping":
			t.Log(rec.Ping(c2[i]))
		}
	}
}
