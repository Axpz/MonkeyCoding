package leetcode

func longestCommonPrefix(names []string) string {
	if len(names) == 0 {
		return ""
	}

	prefix := names[0]
	prefixLen := len(prefix)
	for _, name := range names[1:] {
		j := 0
		for j < prefixLen && j < len(name) && prefix[j] == name[j] {
			j++
		}
		if j == 0 {
			return ""
		}
		prefixLen = j
	}
	return prefix[:prefixLen]
}
