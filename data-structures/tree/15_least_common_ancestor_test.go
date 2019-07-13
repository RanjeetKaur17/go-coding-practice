package tree

import "testing"

/*Problem Statement:
Find the lowest common ancestor in an unordered binary tree given two values in the tree.

Sample Input:
        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2_     0        8
         /   \
         7    4
5 & 1

Sample Output:
3
*/

func Test_LCA(t *testing.T) {
	root := Tree()

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

func Tree() *treeNode {
	return &treeNode{
		value:3,
		left:&treeNode{
			value:5,
			left:&treeNode{
				value:6,
			},
			right:&treeNode{
				value:2,
				left:&treeNode{
					value:7,
				},
				right:&treeNode{
					value:4,
				},
			},
		},
		right:&treeNode{
			value:1,
			left:&treeNode{
				value:0,
			},
			right:&treeNode{
				value:8,
			},
		},
	}
}