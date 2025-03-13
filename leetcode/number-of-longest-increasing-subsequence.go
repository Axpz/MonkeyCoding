package leetcode

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	count := make([]int, n)
	maxLen := 0
	result := 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		count[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
			result = count[i]
		} else if dp[i] == maxLen {
			result += count[i]
		}
	}

	return result
}
