package tree

import (
	"github.com/RanjeetKaur17/go-coding-practice/data-structures/linked_list"
	"testing"
)

func Test_sortedListToBST(t *testing.T) {
	var list *linked_list.ListNode
	bst := sortedListToBST(list)
	if bst != nil {
		t.Fail()
	}

	list = &linked_list.ListNode{1,
		&linked_list.ListNode{2,
			&linked_list.ListNode{3,
				&linked_list.ListNode{4,
					&linked_list.ListNode{5,
						&linked_list.ListNode{6,
							&linked_list.ListNode{7,
								nil}}}}}}}
	bst = sortedListToBST(list)

	if bst.value != 4 || bst.left.value != 2 || bst.left.left.value != 1 || bst.left.right.value != 3 ||
		bst.right.value != 6 || bst.right.left.value != 5 || bst.right.right.value != 7 {
		t.Fail()
	}

	list = &linked_list.ListNode{1,
		&linked_list.ListNode{2,
			&linked_list.ListNode{3,
				&linked_list.ListNode{4,
					&linked_list.ListNode{5,
						&linked_list.ListNode{6,
							nil}}}}}}
	bst = sortedListToBST(list)

	if bst.value != 3 || bst.left.value != 1 || bst.left.right.value != 2 ||
		bst.right.value != 5 || bst.right.left.value != 4 || bst.right.right.value != 6 {
		t.Fail()
	}
}