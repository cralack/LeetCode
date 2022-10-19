package bestsightseeingpair

import "testing"

func maxScoreSightseeingPair(values []int) (ans int) {
	mx := values[0]
	for i := 1; i < len(values); i++ {
		ans = max(ans, values[i]-i+mx)
		mx = max(mx, values[i]+i)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func Test_best_sightseeing_pair(t *testing.T) {
	values := []int{8, 1, 5, 2, 6}
	t.Log(maxScoreSightseeingPair(values))
}
