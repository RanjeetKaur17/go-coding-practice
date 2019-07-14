package tree

import "testing"

func Test_Height(t *testing.T) {
	root := SampleTree()

	height := Height(root)
	if height != 4 {
		t.Fail()
	}

	height = Height(nil)
	if height != 0 {
		t.Fail()
	}

	height = Height(&Tree{value:1})
	if height != 1 {
		t.Fail()
	}
}