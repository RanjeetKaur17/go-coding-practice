package binary_search

import "testing"

func Test_SearchRotatedArray(t *testing.T) {
	i := SearchRotatedArray([]int{0}, 0)

	if i != 0 {
		t.Fail()
	}


	i = SearchRotatedArray([]int{0}, 2)

	if i != -1 {
		t.Fail()
	}

	i = SearchRotatedArray([]int{1,2},2)

	if i != 1 {
		t.Fail()
	}

	i = SearchRotatedArray([]int{2,1}, 2)

	if i != 0 {
		t.Fail()
	}

	i = SearchRotatedArray([]int{3,4,1,2},3)

	if i != 0 {
		t.Fail()
	}

	i = SearchRotatedArray([]int{1,2,3,4},5)

	if i != -1 {
		t.Fail()
	}
}
