package leetcode

// import "fmt"

var count int = 0

func lenLongestFibSubseq(arr []int) int {
	n := len(arr)
	if n < 3 {
		return 0
	}

	// to memeorize the value to reflact the index
	index := map[int]int{}
	for i, v := range arr {
		index[v] = i
	}

	result := 0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			x, y := arr[i], arr[j]
			count := 2
			for true {
				if k, ok := index[x+y]; ok && k > j {
					count += 1
					result = max(result, count)
					x, y = y, x+y
				} else {
					break
				}
			}
		}
	}

	return result
}
