package binary_search

import (
	"testing"
)

func Test_MinRotatedArray(t *testing.T) {
	min := MinRotatedArray([]int{0})

	if min != 0 {
		t.Fail()
	}

	min = MinRotatedArray([]int{1,2})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArray([]int{2,1})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArray([]int{1,2,3,4})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArray([]int{2,3,4,1})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArray([]int{4,1,2,3})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArray([]int{4,5,6,7,0,1,2})

	if min != 0 {
		t.Fail()
	}
}

func Test_MinRotatedArrayIndex(t *testing.T) {
	min := MinRotatedArrayIndex([]int{0})

	if min != 0 {
		t.Fail()
	}

	min = MinRotatedArrayIndex([]int{1,2})

	if min != 0 {
		t.Fail()
	}

	min = MinRotatedArrayIndex([]int{2,1})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArrayIndex([]int{1,2,3,4})

	if min != 0 {
		t.Fail()
	}

	min = MinRotatedArrayIndex([]int{2,3,4,1})

	if min != 3 {
		t.Fail()
	}

	min = MinRotatedArrayIndex([]int{4,1,2,3})

	if min != 1 {
		t.Fail()
	}

	min = MinRotatedArrayIndex([]int{4,5,6,7,0,1,2})

	if min != 4 {
		t.Fail()
	}
}
