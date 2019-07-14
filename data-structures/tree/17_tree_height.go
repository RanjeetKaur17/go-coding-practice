package tree

/*Problem Statement:
Given a binary tree, find its maximum depth.
The maximum depth of a binary tree is the number of nodes along the longest path from the root node down to the farthest leaf node.

Sample Input:
        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2_     0        8
         /   \
         7    4

Sample Output:
4
*/

//Get Height of a binary tree
func Height(A *BinaryTree)  int {
	//If current node is nil, return height as 0
	if A == nil {
		return 0
	}
	//If current node is a leaf node, return height as 1
	if A.left == nil && A.right == nil {
		return 1
	}

	leftHeight := Height(A.left)
	rightHeight := Height(A.right)

	//height of current node = max(height of left subtree, height of right sub tree) + 1
	if leftHeight > rightHeight {
		return leftHeight + 1
	} else {
		return rightHeight + 1
	}
}