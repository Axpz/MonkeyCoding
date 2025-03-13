package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	var result [][]int
	var used []bool = make([]bool, len(nums))
	var dfs func(int, []int)
	dfs = func(depth int, path []int) {
		if depth == len(nums) {
			result = append(result, append([]int{}, path...))
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] || (i > 0 && !used[i-1] && nums[i] == nums[i-1]) {
				continue
			}

			used[i] = true
			path = append(path, nums[i])
			dfs(depth+1, path)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	sort.Ints(nums)
	dfs(0, []int{})
	return result
}
