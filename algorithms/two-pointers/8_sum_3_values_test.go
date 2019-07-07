package two_pointers

import (
	"testing"
)

func Test_SquareRoot(t *testing.T) {
	sum := ThreeSumClosest([]int{0}, 5)

	if sum != 0 {
		t.Fail()
	}

	sum = ThreeSumClosest([]int{0,1}, 5)

	if sum != 1 {
		t.Fail()
	}

	sum = ThreeSumClosest([]int{0,1,2}, 5)

	if sum != 3 {
		t.Fail()
	}

	sum = ThreeSumClosest([]int{-1,2,1,-4}, 1)

	if sum != 2 {
		t.Fail()
	}

	sum = ThreeSumClosest([]int{9, -3, -7, 5, 2, -6, 3, -9, -10, 5, -2, -5, 5, 2, -7, 6, -4, -7, -9, -7, -8, -6, 6, 7, 8, -6, 2, -10, -6, -1, -4, -1, 1, 5, -4, -9, -10, 2, -10, 4, -3, 4, 10, 2, 3}, -2)

	if sum != -2 {
		t.Fail()
	}
}
