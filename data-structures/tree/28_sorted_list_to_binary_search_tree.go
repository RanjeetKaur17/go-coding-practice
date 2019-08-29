package tree

import "github.com/RanjeetKaur17/go-coding-practice/data-structures/linked_list"

func sortedListToBST(list *linked_list.ListNode) *BinaryTreeNode {
	var m []*linked_list.ListNode

	t := list
	for t != nil {
		m = append(m, t)
		t = t.Next
	}
	return convertToBST(m, 0, len(m)-1)
}

func convertToBST(list []*linked_list.ListNode, start int, end int) *BinaryTreeNode {
	if end < start {
		return nil
	} else if start == end {
		return &BinaryTreeNode{value:list[start].Value}
	}

	mid := ((end-start)/2) + start

	return &BinaryTreeNode{value:list[mid].Value, left:convertToBST(list, start, mid-1), right:convertToBST(list, mid+1, end)}
}