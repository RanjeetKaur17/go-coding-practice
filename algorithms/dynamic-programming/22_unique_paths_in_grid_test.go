package dynamic_programming

import "testing"

func Test_UniquePathsWithObstacles(t *testing.T) {

	res := UniquePathsWithObstacles([][]int{{1,0}})
	if res != 0 {
		t.Fail()
	}

	res = UniquePathsWithObstacles([][]int{{0,0,0}, {0,1,0}, {0,0,0}})
	if res != 2 {
		t.Fail()
	}
}
