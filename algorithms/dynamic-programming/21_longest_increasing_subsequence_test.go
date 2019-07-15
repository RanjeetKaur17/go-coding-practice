package dynamic_programming

import "testing"

func TestLongestIncreasingSubsequence(t *testing.T) {
	res := LongestIncreasingSubsequence([]int{})

	if res != 0 {
		t.Fail()
	}

	res = LongestIncreasingSubsequence([]int{1})

	if res != 1 {
		t.Fail()
	}

	res = LongestIncreasingSubsequence([]int{3, 2, 1})

	if res != 1 {
		t.Fail()
	}

	res = LongestIncreasingSubsequence([]int{1, 2, 3})

	if res != 3 {
		t.Fail()
	}

	res = LongestIncreasingSubsequence([]int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15})

	if res != 6 {
		t.Fail()
	}
}
