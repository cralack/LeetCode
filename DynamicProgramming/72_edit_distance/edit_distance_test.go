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
func minDistance_DP(word1, word2 string) int {
	//init
	m, n := len(word1), len(word2)
	if n*m == 0 {
		return m + n
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	//if word2=="" word1->base case
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	//if word1=="" word2->base case
	for j := 1; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+1, min(dp[i][j-1]+1, dp[i-1][j-1]+1))
			}
		}
	}

	return dp[m][n] //O(n^2)
}
func Test_edit_distance(t *testing.T) {
	word1 := "horse"
	word2 := "ros"
	t.Log(minDistance_DFS(word1, word2))

	word1 = "intention"
	word2 = "execution"
	t.Log(minDistance_DP(word1, word2))

	word1 = "dinitrophenylhydrazine"
	word2 = "benzalphenylhydrazone"
	t.Log(minDistance_DP(word1, word2))

	word1 = ""
	word2 = "a"
	t.Log(minDistance_DP(word1, word2))
}

func Benchmark_minDistance(b *testing.B) {
	word1 := "punishment"
	word2 := "production"
	b.Run("DFS", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			minDistance_DFS(word1, word2)
		}
		b.StopTimer()
	})
	b.Run("DP", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			minDistance_DP(word1, word2)
		}
		b.StopTimer()
	})
}
