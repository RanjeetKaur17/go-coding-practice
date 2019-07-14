package tree

import "testing"

func Test_Height(t *testing.T) {
	root := SampleBinaryTree()

	height := Height(root)
	if height != 4 {
		t.Fail()
	}

	height = Height(nil)
	if height != 0 {
		t.Fail()
	}

	height = Height(&BinaryTreeNode{value: 1})
	if height != 1 {
		t.Fail()
	}
}