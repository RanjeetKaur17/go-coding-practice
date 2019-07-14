package tree

import "testing"

func Test_LCA(t *testing.T) {
	root := SampleBinaryTree()

	lca := LCA(root, 7,4)
	if lca != 2 {
		t.Fail()
	}

	lca = LCA(root, 7,2)
	if lca != 2 {
		t.Fail()
	}

	lca = LCA(root, 6,4)
	if lca != 5 {
		t.Fail()
	}

	lca = LCA(root, 6,5)
	if lca != 5 {
		t.Fail()
	}

	lca = LCA(root, 6,1)
	if lca != 3 {
		t.Fail()
	}

	lca = LCA(root, 4,3)
	if lca != 3 {
		t.Fail()
	}

	lca = LCA(root, 4,9)
	if lca != -1 {
		t.Fail()
	}
}