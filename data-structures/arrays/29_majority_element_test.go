package arrays

import "testing"

func Test_MajorityElement(t *testing.T) {
	min := MajorityElement([]int{})

	if min != -1 {
gs
t.Fail()
	}

	min = MajorityElement([]int{1})

	if min != 1 {
		t.Fail()
	}

	min = MajorityElement([]int{2, 1, 2})

	if min != 2 {
		t.Fail()
	}

	min = MajorityElement([]int{2, 1, 1, 3, 1, 1, 2, 4, 5, 1, 1})

	if min != 1 {
		t.Fail()
	}
}