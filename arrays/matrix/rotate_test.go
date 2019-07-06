package matrix

import (
	"testing"
)

func Test_rotate(t *testing.T) {
	A := rotate([][]int{
		{1},
	})

	if !matrxiEquals(A, [][]int{{1}}) {
		t.Fail()
	}

	A = rotate([][]int{
		{1,2},
		{3,4},
	})

	if !matrxiEquals(A, [][]int{{3,1}, {4,2}}) {
		t.Fail()
	}

	A = rotate([][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	})

	if !matrxiEquals(A, [][]int{{7,4,1}, {8,5,2}, {9,6,3}}) {
		t.Fail()
	}

	A = rotate([][]int{
		{1,2,3,4},
		{5,6,7,8},
		{9,10,11,12},
		{13,14,15,16},
	})

	if !matrxiEquals(A, [][]int{{13,9,5,1}, {14,10,6,2}, {15,11,7,3}, {16,12,8,4}}) {
		t.Fail()
	}
}

func matrxiEquals(a [][]int, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if !arrayEquals(v, b[i]) {
			return false
		}
	}
	return true
}

func arrayEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}