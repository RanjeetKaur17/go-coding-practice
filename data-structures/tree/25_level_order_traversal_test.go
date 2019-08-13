package tree

import (
	"testing"
)

func Test_levelOrder(t *testing.T) {
	root := SampleBinaryTree()

	res := levelOrder(root)
	if len(res) != 4 || len(res[0]) != 1 || res[0][0] != 3 ||
		len(res[1]) != 2 || res[1][0] != 5 || res[1][1] != 1 ||
		len(res[2]) != 4 || res[2][0] != 6 || res[2][1] != 2 || res[2][2] != 0 || res[2][3] != 8 ||
		len(res[3]) != 2 || res[3][0] != 7 || res[3][1] != 4 {
		t.Fail()
	}

	root = nil

	res = levelOrder(root)
	if len(res) != 0 {
		t.Fail()
	}

	root = &BinaryTreeNode{value:1}

	res = levelOrder(root)
	if len(res) != 1 || res[0][0] != 1 {
		t.Fail()
	}
}