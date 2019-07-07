package binary_search

import (
	"testing"
)

func Test_Search(t *testing.T) {
	i := Search([]int{0}, 0)

	if i != 0 {
		t.Fail()
	}

	i = Search([]int{0}, 2)

	if i != -1 {
		t.Fail()
	}

	i = Search([]int{1,2},2)

	if i != 1 {
		t.Fail()
	}

	i = Search([]int{1,2}, 1)

	if i != 0 {
		t.Fail()
	}

	i = Search([]int{1,2,3,4},3)

	if i != 2 {
		t.Fail()
	}

	i = Search([]int{1,2,3,4},5)

	if i != -1 {
		t.Fail()
	}
}
