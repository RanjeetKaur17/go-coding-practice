package tree

type BinaryTree struct {
	value int
	left *BinaryTree
	right *BinaryTree
}

func SampleBinaryTree() *BinaryTree {
	return &BinaryTree{
		value:3,
		left:&BinaryTree{
			value:5,
			left:&BinaryTree{
				value:6,
			},
			right:&BinaryTree{
				value:2,
				left:&BinaryTree{
					value:7,
				},
				right:&BinaryTree{
					value:4,
				},
			},
		},
		right:&BinaryTree{
			value:1,
			left:&BinaryTree{
				value:0,
			},
			right:&BinaryTree{
				value:8,
			},
		},
	}
}

func SampleBST() *BinaryTree {
	return &BinaryTree{
		value:10,
		left:&BinaryTree{
			value:5,
			left:&BinaryTree{
				value:2,
			},
			right:&BinaryTree{
				value:7,
				left:&BinaryTree{
					value:6,
				},
				right:&BinaryTree{
					value:8,
				},
			},
		},
		right:&BinaryTree{
			value:15,
			left:&BinaryTree{
				value:12,
			},
			right:&BinaryTree{
				value:20,
			},
		},
	}
}