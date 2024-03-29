package arrays

func MatrxEquals(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !ArrayEquals(v, b[i]) {
			return false
		}
	}
	return true
}

func ArrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}