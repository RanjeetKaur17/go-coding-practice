package binary_search

import (
	"testing"
)

func Test_MinRotatedMatrixt(t *testing.T) {
	min := MinRotatedMatrix([]int{0})

	if min != 0 {
		t.Fail()
	}

	min = MinRotatedMatrix([]int{1,2})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedMatrix([]int{2,1})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedMatrix([]int{1,2,3,4})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedMatrix([]int{2,3,4,1})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedMatrix([]int{4,1,2,3})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedMatrix([]int{4,5,6,7,0,1,2})

	if min != 0 {
		t.Fail()
	}
}
