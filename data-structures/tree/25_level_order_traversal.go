package tree

/*
Problem Statement:
Given a binary tree, return the level order traversal of its nodes’ values. (ie, from left to right, level by level).

Sample Input:
Sample Input:
        _______3______
       /              \
    ___5__          ___1__
   /      \        /      \
   6      _2_     0        8
         /   \
         7    4

Sample Output:
[
  [3],
  [5,1],
  [6,2,0,8],
  [7,4]
]
*/

//returns level order traversal
func levelOrder(A *BinaryTreeNode )  [][]int {
	if A == nil {
		return nil
	}
	c := make(chan interface{}, 100)
	h := 0
	c <- A
	c <- h
	var res [][]int
	res = append(res, []int{})
	for len(c) != 0 {
		v := <- c
		switch v.(type) {
		case int:
			if len(c) != 0 {
				h++
				c <- h
				res = append(res, []int{})
			}
		case *BinaryTreeNode:
			vt := v.(*BinaryTreeNode)
			res[h] = append(res[h], vt.value)
			if vt.left != nil {
				c <- vt.left
			}

			if vt.right != nil {
				c <- vt.right
			}
		}
	}

	return res
}