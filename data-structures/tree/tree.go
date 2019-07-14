package tree

type BinaryTreeNode struct {
	value int
	left *BinaryTreeNode
	right *BinaryTreeNode
}

func SampleBinaryTree() *BinaryTreeNode {
	return &BinaryTreeNode{
		value:3,
		left:&BinaryTreeNode{
			value:5,
			left:&BinaryTreeNode{
				value:6,
			},
			right:&BinaryTreeNode{
				value:2,
				left:&BinaryTreeNode{
					value:7,
				},
				right:&BinaryTreeNode{
					value:4,
				},
			},
		},
		right:&BinaryTreeNode{
			value:1,
			left:&BinaryTreeNode{
				value:0,
			},
			right:&BinaryTreeNode{
				value:8,
			},
		},
	}
}

func SampleBST() *BinaryTreeNode {
	return &BinaryTreeNode{
		value:10,
		left:&BinaryTreeNode{
			value:5,
			left:&BinaryTreeNode{
				value:2,
			},
			right:&BinaryTreeNode{
				value:7,
				left:&BinaryTreeNode{
					value:6,
				},
				right:&BinaryTreeNode{
					value:8,
				},
			},
		},
		right:&BinaryTreeNode{
			value:15,
			left:&BinaryTreeNode{
				value:12,
			},
			right:&BinaryTreeNode{
				value:20,
			},
		},
	}
}