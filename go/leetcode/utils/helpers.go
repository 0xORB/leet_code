package utils

// Capitalize function name to make it public
func EqualSlice[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func EqualSliceIgnoreOrder[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	countMap := make(map[T]int)
	for _, val := range a {
		countMap[val]++
	}

	for _, val := range b {
		countMap[val]--
		if countMap[val] < 0 {
			return false
		}
	}

	for _, count := range countMap {
		if count != 0 {
			return false
		}
	}

	return true
}
