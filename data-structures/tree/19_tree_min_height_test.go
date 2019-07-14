package tree

import "testing"

func Test_MinHeight(t *testing.T) {
	root := SampleBinaryTree()

	height := MinHeight(root)
	if height != 3 {
		t.Fail()
	}

	height = MinHeight(nil)
	if height != 0 {
		t.Fail()
	}

	height = MinHeight(&BinaryTreeNode{value: 1})
	if height != 1 {
		t.Fail()
	}
}