package two_pointers

import (
	"testing"
)

func Test_DiffKPossible(t *testing.T) {
	i := DiffKPossible([]int{0}, 0)

	if i != 0 {
		t.Fail()
	}

	i = DiffKPossible([]int{0}, 2)

	if i != 0 {
		t.Fail()
	}

	i = DiffKPossible([]int{1,2},2)

	if i != 0 {
		t.Fail()
	}

	i = DiffKPossible([]int{1,2}, 1)

	if i != 1 {
		t.Fail()
	}

	i = DiffKPossible([]int{1,2,3,4},3)

	if i != 1 {
		t.Fail()
	}

	i = DiffKPossible([]int{1,2,3,4},5)

	if i != 0 {
		t.Fail()
	}

	i = DiffKPossible([]int{1,2,3,4,8},6)

	if i != 1 {
		t.Fail()
	}
}
