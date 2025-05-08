package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	m := make(map[rune]int, len(ransomNote))

	for _, v := range magazine {
		m[v] += 1
	}

	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		} else {
			m[v]--
		}
	}
	return true
}
