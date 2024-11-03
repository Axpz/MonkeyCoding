package leetcode

import (
	"sort"
)

func repeatLimitedString(s string, repeatLimit int) string {
	arrs := []rune(s)
	m := map[rune]int{}
	sList := []rune{}
	for _, r := range arrs {
		if _, ok := m[r]; !ok {
			sList = append(sList, r)
		}

		m[r]++
	}

	sort.Slice(sList, func(i, j int) bool {
		return sList[i] < sList[j]
	})

	// z...a
	var result []rune
	i := len(sList) - 1
	for i >= 0 {
		r := sList[i]
		j := 0

		for j < repeatLimit && m[r] > 0 {
			result = append(result, r)
			j++
			m[r]--
		}

		if m[r] <= 0 {
			i--
			continue
		}

		k := i - 1
		for k >= 0 && m[sList[k]] <= 0 {
			k--
		}
		if k >= 0 {
			// result += string(sList[k])
			result = append(result, sList[k])
			m[sList[k]]--
		} else {
			break
		}
	}

	return string(result)
}
