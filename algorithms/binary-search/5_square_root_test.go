package binary_search

import (
	"testing"
)

func Test_SquareRoot(t *testing.T) {
	root := SquareRoot(0)

	if root != 0 {
		t.Fail()
	}

	root = SquareRoot(1)

	if root != 1 {
		t.Fail()
	}

	root = SquareRoot(2)

	if root != 1 {
		t.Fail()
	}

	root = SquareRoot(4)

	if root != 2 {
		t.Fail()
	}

	root = SquareRoot(11)

	if root != 3 {
		t.Fail()
	}

	root = SquareRoot(81)

	if root != 9 {
		t.Fail()
	}

	root = SquareRoot(-4)

	if root != 0 {
		t.Fail()
	}
}
