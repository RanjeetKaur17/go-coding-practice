package tree

import (
	"testing"
)

func Test_BSTIterator(t *testing.T) {
	iter := NewBSTIterator(SampleBST())

	prev := 0
	count := 0

	for iter.hasNext() {
		if iter.next() < prev {
			t.Fail()
		}
		count++
	}

	if count != 9 {
		t.Fail()
	}

	//single value tree
	iter = NewBSTIterator(&BinaryTreeNode{value:1})

	if !iter.hasNext() {
		t.Fail()
	}
	if iter.next() != 1 {
		t.Fail()
	}
	if iter.hasNext() {
		t.Fail()
	}

	//empty tree
	iter = NewBSTIterator(nil)

	if iter.hasNext() {
		t.Fail()
	}
}