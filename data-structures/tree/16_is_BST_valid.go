package tree

/*Problem Statement:
Given a binary tree, determine if it is a valid binary search tree (BST).

Sample Input:
    1
  /  \
 2    3

SampleOutput:
False
*/

//Check if given tree is a valid BST(binary search tree)
func IsValidBST(A *BinaryTree)  bool {
	//get in-order traversal
	a := inOrder(A)

	//check if inorder traversal is sorted or not
	for i := 1; i < len(a); i++ {
		if a[i] <= a[i-1] {
			//return false if unsorted
			return false
		}
	}
	//return true if sorted
	return true
}

//Get InOrder traversal of a tree
func inOrder(A *BinaryTree) []int {
	//if current node is nil return nothing
	if A == nil {
		return  nil
	}
	//if current node is a leaf node, return single value array
	if A.left == nil && A.right == nil {
		return []int{A.value}
	}

	//inOrder traversal for left subtree + current node + inOrder traversal for right sub tree = inOrder traversal for current node
	return append(append(inOrder(A.left), A.value), inOrder(A.right)...)
}