package tree

import "testing"

func Test_LCA(t *testing.T) {
	root := SampleTree()

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

func SampleTree() *Tree {
	return &Tree{
		value:3,
		left:&Tree{
			value:5,
			left:&Tree{
				value:6,
			},
			right:&Tree{
				value:2,
				left:&Tree{
					value:7,
				},
				right:&Tree{
					value:4,
				},
			},
		},
		right:&Tree{
			value:1,
			left:&Tree{
				value:0,
			},
			right:&Tree{
				value:8,
			},
		},
	}
}