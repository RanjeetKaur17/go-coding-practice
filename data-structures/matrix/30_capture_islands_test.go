package matrix

import (
	"testing"
)

func Test_CaptureRegions(t *testing.T) {
	res := CaptureRegions([][]rune{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	})

	if res[1][1] != 'X' || res[1][2] != 'X' || res[2][2] != 'X' || res[3][1] != 'O' || res[4][1] != 'O' {
		t.Fail()
	}
}