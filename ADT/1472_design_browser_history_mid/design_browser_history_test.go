package design_browser_history_mid

import (
	"testing"
)

type BrowserHistory struct {
	stk []string
	cur int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		stk: []string{homepage},
		cur: 0,
	}
}

func (h *BrowserHistory) Visit(url string) {
	h.stk = h.stk[:h.cur+1]
	h.stk = append(h.stk, url)
	h.cur++
}

func (h *BrowserHistory) Back(steps int) string {
	h.cur = max(0, h.cur-steps)
	return h.stk[h.cur]
}

func (h *BrowserHistory) Forward(steps int) string {
	h.cur = min(h.cur+steps, len(h.stk)-1)
	return h.stk[h.cur]
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

func Test_design_browser_history(t *testing.T) {
	tests := []struct {
		cmd1 []string
		cmd2 []interface{}
	}{
		{
			cmd1: []string{"BrowserHistory", "visit", "visit", "visit", "back", "back",
				"forward", "visit", "forward", "back", "back"},
			cmd2: []interface{}{"leetcode.com", "google.com", "facebook.com", "youtube.com",
				1, 1, 1, "linkedin.com", 2, 2, 7},
		},
	}
	for _, tt := range tests {
		var sol BrowserHistory
		for i, c := range tt.cmd1 {
			switch c {
			case "BrowserHistory":
				sol = Constructor(tt.cmd2[i].(string))
			case "visit":
				sol.Visit(tt.cmd2[i].(string))
			case "back":
				t.Log(sol.Back(tt.cmd2[i].(int)))
			case "forward":
				t.Log(sol.Forward(tt.cmd2[i].(int)))
			}
		}
	}
}
