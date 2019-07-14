package tree

/*Problem Statement:
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.
The first call to next() will return the smallest number in BST. Calling next() again will return the next smallest number in the BST, and so on.
*/

//Iterator for binary search tree
type BSTIterator struct {
	tree  *BinaryTreeNode;
	stack []*BinaryTreeNode;
}

//Create new iterator object
func NewBSTIterator(root *BinaryTreeNode) BSTIterator {
	iter := BSTIterator{tree:root}

	//add all nodes along left path of root, into a stack
	for root != nil {
		iter.stack = append(iter.stack, root);
		root = root.left;
	}
	return iter
}

//check if BST has next values
func (iter *BSTIterator) hasNext() bool {
	if len(iter.stack) == 0 {
		return false;
	}
	return true;
}

//get next value from BST
func (iter *BSTIterator) next() int{
	if len(iter.stack) == 0 {
		return 0;
	}
	// current value of iterator is top of stack
	c := iter.stack[len(iter.stack)-1];
	iter.stack = iter.stack[:len(iter.stack)-1]

	//add all nodes along left path of right subtree of current node, into the stack
	t := c;
	if t.right != nil {
		t = t.right;
		for t != nil {
			iter.stack = append(iter.stack, t);
			t = t.left;
		}
	}

	return c.value;
}