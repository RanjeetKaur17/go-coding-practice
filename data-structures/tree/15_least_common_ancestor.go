package tree

/*
Problem Statement:
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

func LCA(A *BinaryTreeNode, B int , C int )  (int) {
	l, _, _ := lca(A, B, C)
	return l
}

func lca(A *BinaryTreeNode, B int , C int )  (int, bool, bool) {
	if A == nil {
		return -1, false, false
	} else {
		//If left and Right both found in left subtree, return solution
		l, bFoundInLeft, cFoundinLeft := lca(A.left, B, C)
		if bFoundInLeft && cFoundinLeft {
			return l, bFoundInLeft, cFoundinLeft
		}

		//If left and Right both found in right subtree, return solution
		r, bFoundInRight, cFoundInRight := lca(A.right, B, C)
		if bFoundInRight && cFoundInRight {
			return r, bFoundInRight, cFoundInRight
		}

		//If current element is the LCA, return current value as solution
		if ((bFoundInLeft && cFoundInRight) || (bFoundInRight && cFoundinLeft)) ||
			((bFoundInLeft || bFoundInRight) && A.value == C) ||
			((cFoundinLeft || cFoundInRight) && A.value == B) {
			return A.value, true, true
		}

		//If either value found in left subtree, return that information
		if bFoundInLeft || cFoundinLeft {
			return l, bFoundInLeft, cFoundinLeft
		}

		//If either value found in right subtree, return that information
		if bFoundInRight || cFoundInRight {
			return r, bFoundInRight, cFoundInRight
		}

		//If current value equals either of the values, return that information
		if A.value == B {
			return -1, true, false
		}
		if A.value == C {
			return -1, false, true
		}

		//return negative result
		return -1, false, false
	}
}