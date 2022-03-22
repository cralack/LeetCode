package editdistance

import (
	"testing"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDistance_DFS(word1 string, word2 string) int {

	var dp func(i, j int) int
	dp = func(i, j int) int {
		//base case
		if i == -1 {
			return j + 1
		}
		if j == -1 {
			return i + 1
		}
		//skip
		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		} else {
			return min(dp(i, j-1)+1, min(dp(i-1, j)+1, dp(i-1, j-1)+1))
		}
	}
	return dp(len(word1)-1, len(word2)-1)
}
func minDistance_MEMO(s1, s2 string) int {
	//init
	m, n := len(s1), len(s2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
		if i >= 1 {
			dp[i][0] = i
		}
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			c1, c2 := string(s1[i-1]), string(s2[j-1])
			if c1 == c2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}

	return dp[m][n]
}
func Test_edit_distance(t *testing.T) {
	s1 := "horse"
	s2 := "ros"
	t.Log(minDistance_DFS(s1, s2))

	s1 = "intention"
	s2 = "execution"
	t.Log(minDistance_MEMO(s1, s2))

	s1 = "dinitrophenylhydrazine"
	s2 = "benzalphenylhydrazone"
	t.Log(minDistance_MEMO(s1, s2))
}
