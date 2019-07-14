package tree

/*Problem Statement:
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Sample Input:
        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2_     0        8
         /   \
         7    4

Sample Output:
3
*/

//Get Minimum Height of a binary tree
func MinHeight(A *BinaryTreeNode)  int {
	//If current node is nil, return height as 0
	if A == nil {
		return 0
	}
	//If current node is a leaf node, return height as 1
	if A.left == nil && A.right == nil {
		return 1
	}

	leftHeight := MinHeight(A.left)
	//If node has only left node, return height as per left node
	if A.right == nil {
		return leftHeight + 1
	}
	rightHeight := MinHeight(A.right)
	//If node has only right node, return height as per right node
	if A.left == nil {
		return rightHeight + 1
	}
	//height of current node = min(height of left subtree, height of right sub tree) + 1
	if leftHeight < rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}