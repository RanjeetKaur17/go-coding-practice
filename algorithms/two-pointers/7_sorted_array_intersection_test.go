package two_pointers

import (
	"testing"
)

func Test_Intersect(t *testing.T) {
	C := Intersect([]int{0}, []int{0})

	if len(C) != 1 || C[0] != 0 {
		t.Fail()
	}

	C = Intersect([]int{0}, []int{2})

	if len(C) != 0 {
		t.Fail()
	}

	C = Intersect([]int{1,2},[]int{0})

	if len(C) != 0 {
		t.Fail()
	}

	C = Intersect([]int{1,2}, []int{1})

	if len(C) != 1 || C[0] != 1 {
		t.Fail()
	}

	C = Intersect([]int{1,2,3,4},[]int{1,2})

	if len(C) != 2 || C[0] != 1 || C[1] != 2 {
		t.Fail()
	}

	C = Intersect([]int{1,2},[]int{1,2})

	if len(C) != 2 || C[0] != 1 || C[1] != 2 {
		t.Fail()
	}
}
