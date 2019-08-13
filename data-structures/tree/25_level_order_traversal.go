package tree

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