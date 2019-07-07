package matrix

import (
	"github.com/RanjeetKaur17/go-coding-practice/data-structures/arrays"
	"testing"
)

func Test_rotate(t *testing.T) {
	A := rotate([][]int{
		{1},
	})

	if !arrays.MatrxEquals(A, [][]int{{1}}) {
		t.Fail()
	}

	A = rotate([][]int{
		{1,2},
		{3,4},
	})

	if !arrays.MatrxEquals(A, [][]int{{3,1}, {4,2}}) {
		t.Fail()
	}

	A = rotate([][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	})

	if !arrays.MatrxEquals(A, [][]int{{7,4,1}, {8,5,2}, {9,6,3}}) {
		t.Fail()
	}

	A = rotate([][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	})

	if !arrays.MatrxEquals(A, [][]int{{13,9,5,1}, {14,10,6,2}, {15,11,7,3}, {16,12,8,4}}) {
		t.Fail()
	}
}