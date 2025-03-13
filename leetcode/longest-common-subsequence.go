package leetcode

func longestCommonSubsequence(text1, text2 string) int {
	t1, t2 := " "+text1, " "+text2
	m, n := len(t1), len(t2)

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if t1[i] == t2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}
